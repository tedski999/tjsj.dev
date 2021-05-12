package webstats

import "time"

// Return the total number of recorded hits
func (stats *Statistics) GetTotalHits() int {
	return stats.totalHits
}

// Return the time since this stats object was started
func (stats *Statistics) GetUptime() string {
	return time.Now().Sub(stats.startTime).Round(time.Millisecond).String()
}
