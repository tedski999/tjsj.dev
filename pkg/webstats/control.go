package webstats

import "time"

// Start the statistics recording
func (stats *Statistics) Start(errChan chan<- error) {
	stats.errChan = errChan
	go stats.Load()
	stats.startTime = time.Now()
}

// Gracefully shutdown the statistics recording
func (stats *Statistics) Stop() {
	stats.Save()
}
