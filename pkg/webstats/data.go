package webstats

import (
	"time"
	"net/http"
	"net/url"
	"sort"
	"strings"
)

// Record data from a new request
func (stats *Statistics) RecordRequest(r *http.Request) {

	// Add to appropriate hit counter
	path := strings.TrimSuffix(r.URL.Path, "/")
	if len(path) == 0 { path = "/" }
	stats.hitCounters[path]++;

	// Add to appropriate referrer counter
	if url, err := url.Parse(r.Referer()); err == nil {
		hostname := url.Hostname()
		if len(hostname) == 0 { hostname = "<direct>" }
		stats.referrerCounters[hostname]++;
	}
}

// Return an list of the top pages by hits
func (stats *Statistics) GetHitCounters() (map[string]int, []string) {
	return stats.hitCounters, sortStringIntMapByValue(stats.hitCounters)
}

// Return an list of the top referrers by referees
func (stats *Statistics) GetReferrerCounters() (map[string]int, []string) {
	return stats.referrerCounters, sortStringIntMapByValue(stats.referrerCounters)
}

// Return the time since this stats object was started
func (stats *Statistics) GetUptime() string {
	return time.Now().Sub(stats.startTime).Round(time.Millisecond).String()
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
