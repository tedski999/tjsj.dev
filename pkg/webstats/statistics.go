package webstats

import "time"

type Statistics struct {
	dataFilePath string
	totalHits int
	startTime time.Time
	errChan chan<- error
}

func Create(dataFilePath string) (*Statistics, error) {

	// Setup statistics
	stats := &Statistics {
		dataFilePath: dataFilePath,
	}

	return stats, nil
}
