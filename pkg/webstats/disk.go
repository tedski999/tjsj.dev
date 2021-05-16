package webstats

import (
	"os"
	"encoding/gob"
	"path/filepath"
)

type statsFileData struct {
	HitCounters map[string]int
	ReferrerCounters map[string]int
	ResponseCodeCounters map[int]int
	TotalUncompressedDataTransferred uint64
	TotalCompressedDataTransferred uint64
}

// Attempt to load stats from disk
func (stats *Statistics) Load() {

	// Open the file
	dataFile, err := os.Open(stats.dataFilePath)
	if err != nil  {
		stats.errChan <- err
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
	stats.hitCounters = data.HitCounters
	stats.referrerCounters = data.ReferrerCounters
	stats.responseCodeCounters = data.ResponseCodeCounters
	stats.totalUncompressedDataTransferred = data.TotalUncompressedDataTransferred
	stats.totalCompressedDataTransferred = data.TotalCompressedDataTransferred
}

// Attempt to save stats to disk
func (stats *Statistics) Save() {

	// Convert stats to data
	data := statsFileData {
		HitCounters: stats.hitCounters,
		ReferrerCounters: stats.referrerCounters,
		ResponseCodeCounters: stats.responseCodeCounters,
		TotalUncompressedDataTransferred: stats.totalUncompressedDataTransferred,
		TotalCompressedDataTransferred: stats.totalCompressedDataTransferred,
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
