package webstats

// Add one to the total hits counter
func (stats *Statistics) IncrementHitCounter() {
	stats.totalHits++
}
