package webstats

import (
	"os"
	"encoding/gob"
	"path/filepath"
)

type statsFileData struct {
	TotalHits int
}

// Attempt to load stats from disk
func (stats *Statistics) Load() {

	// Open the file
	dataFile, err := os.Open(stats.dataFilePath)
	if err != nil  {
		if !os.IsNotExist(err) {
			stats.errChan <- err
		}
		return
	}
	defer dataFile.Close()

	// Load data from the file
	dec := gob.NewDecoder(dataFile)
	var data statsFileData
	if err := dec.Decode(&data); err != nil {
		stats.errChan <- err
		return
	}

	// Convert data to stats
	stats.totalHits = data.TotalHits
}

// Attempt to save stats to disk
func (stats *Statistics) Save() {

	// Convert stats to data
	data := statsFileData {
		TotalHits: stats.totalHits,
	}

	// Create the file
	os.MkdirAll(filepath.Dir(stats.dataFilePath), 0700)
	dataFile, err := os.Create(stats.dataFilePath)
	if err != nil {
		stats.errChan <- err
		return
	}
	defer dataFile.Close()

	// Save data to the file
	enc := gob.NewEncoder(dataFile)
	if err := enc.Encode(data); err != nil {
		stats.errChan <- err
		return
	}
}
