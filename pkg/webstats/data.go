package webstats

import (
	"os"
	"time"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"runtime"
)

type SystemStats struct {
	Hostname, OS, Arch, Uptime string
	CPUCount, CPUUsage int
	GoroutineCount int
	RAMAvailable, RAMUsage uint64
}

// Record data from a new request
func (stats *Statistics) RecordData(w *StatsResponseWriter, r *http.Request) {
	stats.recordDataMutex.Lock()
	defer stats.recordDataMutex.Unlock()

	// Add to appropriate hit counter if the response code was OK
	if w.status == http.StatusOK {
		path := strings.TrimSuffix(r.URL.Path, "/")
		if len(path) == 0 { path = "/" }
		stats.hitCounters[path]++;
	}

	// Add to appropriate referrer counter
	if url, err := url.Parse(r.Referer()); err == nil {
		hostname := url.Hostname()
		if len(hostname) == 0 { hostname = "<direct>" }
		stats.referrerCounters[hostname]++;
	}

	// Record response data
	stats.responseCodeCounters[w.status]++
	stats.totalUncompressedDataTransferred += uint64(w.length)
}

// Record data after compression from a new request
func (stats *Statistics) RecordCompressedData(w *StatsResponseWriter, r *http.Request) {
	stats.recordDataMutex.Lock()
	defer stats.recordDataMutex.Unlock()
	stats.totalCompressedDataTransferred += uint64(w.length)
}

// Return an list of the top pages by hits
func (stats *Statistics) GetHitCounters() (map[string]int, []string) {
	return stats.hitCounters, sortStringIntMapByValue(stats.hitCounters)
}

// Return an list of the top referrers by referees
func (stats *Statistics) GetReferrerCounters() (map[string]int, []string) {
	return stats.referrerCounters, sortStringIntMapByValue(stats.referrerCounters)
}

// Return an list of the top response codes by number of responses
func (stats *Statistics) GetResponseCodeCounters() (map[int]int, []int) {
	return stats.responseCodeCounters, sortIntIntMapByValue(stats.responseCodeCounters)
}

// Return the time since this stats object was started
func (stats *Statistics) GetUptime() string {
	return time.Now().Sub(stats.startTime).Round(time.Millisecond).String()
}

// Return the total number of bytes responded with
func (stats *Statistics) GetTotalDataTransferred() (uint64, uint64) {
	return stats.totalCompressedDataTransferred, stats.totalUncompressedDataTransferred
}

// Return statistics of current system
func (stats *Statistics) GetSystemStats() SystemStats {

	// Hostname
	hostname, err := os.Hostname()
	if err != nil {
		hostname = "<Unknown>"
	}

	// Memory
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)

	// TODO: uptime and cpu usage
	return SystemStats {
		Hostname: hostname,
		OS: runtime.GOOS,
		Arch: runtime.GOARCH,
		Uptime: "<Unknown>",
		CPUCount: runtime.NumCPU(),
		RAMAvailable: memStats.Sys,
		CPUUsage: 0,
		RAMUsage: memStats.Alloc,
		GoroutineCount: runtime.NumGoroutine(),
	}
}

// Return map m keys ordered by map m values
func sortStringIntMapByValue(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys
}

// Return map m keys ordered by map m values
func sortIntIntMapByValue(m map[int]int) []int {
	keys := make([]int, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	return keys
}
