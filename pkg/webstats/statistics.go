package webstats

import "time"

type Statistics struct {
	dataFilePath string
	hitCounters map[string]int
	referrerCounters map[string]int
	startTime time.Time
	errChan chan<- error
}

func Create(dataFilePath string) (*Statistics, error) {

	// Setup statistics
	stats := &Statistics {
		dataFilePath: dataFilePath,
		hitCounters: make(map[string]int),
		referrerCounters: make(map[string]int),
	}

	return stats, nil
}
