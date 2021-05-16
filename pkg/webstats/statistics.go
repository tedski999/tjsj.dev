package webstats

import (
	"time"
	"sync"
)

type Statistics struct {
	dataFilePath string
	hitCounters map[string]int
	referrerCounters map[string]int
	responseCodeCounters map[int]int
	totalUncompressedDataTransferred uint64
	totalCompressedDataTransferred uint64
	startTime time.Time
	recordDataMutex sync.Mutex
	errChan chan<- error
}

func Create(dataFilePath string) (*Statistics, error) {

	// Setup statistics
	stats := &Statistics {
		dataFilePath: dataFilePath,
		hitCounters: make(map[string]int),
		referrerCounters: make(map[string]int),
		responseCodeCounters: make(map[int]int),
		recordDataMutex: sync.Mutex{},
	}

	return stats, nil
}
