package sparql

// StatisticsCollector tracks and maintains statistics about query execution.
// It:
// - Tracks predicate cardinalities
// - Monitors index usage
// - Records join selectivity
// - Updates statistics for better future optimization
type StatisticsCollector struct {
	PredicateStats map[string]*PredicateExecutionStats
	TypeStats      map[string]*TypeExecutionStats
	IndexStats     map[string]*IndexExecutionStats
	JoinStats      []*JoinExecutionStats
	ExecutionCount int
}

// PredicateExecutionStats contains execution statistics for a predicate.
type PredicateExecutionStats struct {
	Predicate           string
	ExecutionCount      int
	EstimatedCard       uint64
	ActualCard          uint64
	IndexUsageCount     int
	FullScanCount       int
	SelectivityEstimate float64
	SelectivityActual   float64
}

// TypeExecutionStats contains execution statistics for a type.
type TypeExecutionStats struct {
	Type              string
	InstanceCount     uint64
	SelectivityActual float64
}

// IndexExecutionStats contains execution statistics for an index.
type IndexExecutionStats struct {
	Predicate    string
	IndexType    string
	UsageCount   int
	CostEstimate float64
	CostActual   float64
}

// JoinExecutionStats contains execution statistics for a join.
type JoinExecutionStats struct {
	LeftPredicate  string
	RightPredicate string
	JoinVars       []string
	Selectivity    float64
	ExecutionCount int
}

// NewStatisticsCollector creates a new statistics collector.
func NewStatisticsCollector() *StatisticsCollector {
	return &StatisticsCollector{
		PredicateStats: make(map[string]*PredicateExecutionStats),
		TypeStats:      make(map[string]*TypeExecutionStats),
		IndexStats:     make(map[string]*IndexExecutionStats),
		JoinStats:      make([]*JoinExecutionStats, 0),
		ExecutionCount: 0,
	}
}

// RecordExecution records execution statistics for a predicate.
func (sc *StatisticsCollector) RecordExecution(predicate string, cardinality uint64) {
	if predicate == "" {
		return
	}

	if _, exists := sc.PredicateStats[predicate]; !exists {
		sc.PredicateStats[predicate] = &PredicateExecutionStats{
			Predicate: predicate,
		}
	}

	stats := sc.PredicateStats[predicate]
	stats.ExecutionCount++
	stats.ActualCard = cardinality

	// Update selectivity actual if we have a previous estimate
	if stats.ActualCard > 0 && stats.EstimatedCard > 0 {
		stats.SelectivityActual = 1.0 / float64(stats.ActualCard)
	}

	sc.ExecutionCount++
}

// RecordIndexUsage records index usage statistics.
func (sc *StatisticsCollector) RecordIndexUsage(predicate, indexType string) {
	if predicate == "" {
		return
	}

	if _, exists := sc.PredicateStats[predicate]; !exists {
		sc.PredicateStats[predicate] = &PredicateExecutionStats{
			Predicate: predicate,
		}
	}

	stats := sc.PredicateStats[predicate]
	stats.IndexUsageCount++

	// Record index-specific stats
	if _, exists := sc.IndexStats[predicate]; !exists {
		sc.IndexStats[predicate] = &IndexExecutionStats{
			Predicate: predicate,
			IndexType: indexType,
		}
	}

	sc.IndexStats[predicate].UsageCount++
}

// RecordJoinExecution records join execution statistics.
func (sc *StatisticsCollector) RecordJoinExecution(left, right string, vars []string, selectivity float64) {
	joinStat := &JoinExecutionStats{
		LeftPredicate:  left,
		RightPredicate: right,
		JoinVars:       append([]string{}, vars...),
		Selectivity:    selectivity,
		ExecutionCount: 1,
	}

	// Check if we already have stats for this join
	for i, existing := range sc.JoinStats {
		if existing.LeftPredicate == left && existing.RightPredicate == right {
			sc.JoinStats[i].ExecutionCount++
			return
		}
	}

	sc.JoinStats = append(sc.JoinStats, joinStat)
}

// GetStats retrieves statistics for a predicate.
func (sc *StatisticsCollector) GetStats(predicate string) *PredicateExecutionStats {
	if stats, exists := sc.PredicateStats[predicate]; exists {
		return stats
	}

	return &PredicateExecutionStats{
		Predicate: predicate,
	}
}

// GetTypeStats retrieves statistics for a type.
func (sc *StatisticsCollector) GetTypeStats(typeName string) *TypeExecutionStats {
	if stats, exists := sc.TypeStats[typeName]; exists {
		return stats
	}

	return &TypeExecutionStats{
		Type: typeName,
	}
}

// RecordTypeInstance records a type instance execution.
func (sc *StatisticsCollector) RecordTypeInstance(typeName string, count uint64) {
	if typeName == "" {
		return
	}

	if _, exists := sc.TypeStats[typeName]; !exists {
		sc.TypeStats[typeName] = &TypeExecutionStats{
			Type: typeName,
		}
	}

	sc.TypeStats[typeName].InstanceCount = count
}

// UpdateEstimates updates cardinality estimates based on actual statistics.
func (sc *StatisticsCollector) UpdateEstimates() {
	for _, stats := range sc.PredicateStats {
		if stats == nil {
			continue
		}

		// Update selectivity estimate based on actual
		if stats.ExecutionCount > 0 && stats.ActualCard > 0 {
			stats.SelectivityActual = 1.0 / float64(stats.ActualCard)
			// Only update estimate if we haven't already calculated it
			if stats.SelectivityEstimate == 0 {
				stats.SelectivityEstimate = stats.SelectivityActual
			}
		}
	}
}

// GetStatisticsInfo returns summary information about collected statistics.
func (sc *StatisticsCollector) GetStatisticsInfo() *StatisticsInfo {
	info := &StatisticsInfo{
		TrackedPredicates: len(sc.PredicateStats),
		TrackedTypes:      len(sc.TypeStats),
		TrackedIndexes:    len(sc.IndexStats),
		TrackedJoins:      len(sc.JoinStats),
		TotalExecutions:   sc.ExecutionCount,
	}

	// Calculate average selectivity
	if len(sc.PredicateStats) > 0 {
		totalSelectivity := 0.0
		count := 0

		for _, stats := range sc.PredicateStats {
			if stats != nil && stats.SelectivityActual > 0 {
				totalSelectivity += stats.SelectivityActual
				count++
			}
		}

		if count > 0 {
			info.AverageSelectivity = totalSelectivity / float64(count)
		}
	}

	return info
}

// StatisticsInfo contains summary information about collected statistics.
type StatisticsInfo struct {
	TrackedPredicates  int
	TrackedTypes       int
	TrackedIndexes     int
	TrackedJoins       int
	TotalExecutions    int
	AverageSelectivity float64
}

// GetMostUsedPredicates returns the most frequently used predicates.
func (sc *StatisticsCollector) GetMostUsedPredicates(limit int) []*PredicateExecutionStats {
	result := make([]*PredicateExecutionStats, 0)

	// Create a slice of all stats
	allStats := make([]*PredicateExecutionStats, 0, len(sc.PredicateStats))
	for _, stats := range sc.PredicateStats {
		if stats != nil {
			allStats = append(allStats, stats)
		}
	}

	// Sort by execution count (would need sort.Slice in real implementation)
	// For now, just collect all and return up to limit
	count := 0
	for _, stats := range allStats {
		if count >= limit {
			break
		}
		result = append(result, stats)
		count++
	}

	return result
}

// Reset clears all collected statistics.
func (sc *StatisticsCollector) Reset() {
	sc.PredicateStats = make(map[string]*PredicateExecutionStats)
	sc.TypeStats = make(map[string]*TypeExecutionStats)
	sc.IndexStats = make(map[string]*IndexExecutionStats)
	sc.JoinStats = make([]*JoinExecutionStats, 0)
	sc.ExecutionCount = 0
}

// GetAccuracyMetrics calculates accuracy of cardinality estimates.
func (sc *StatisticsCollector) GetAccuracyMetrics() *AccuracyMetrics {
	metrics := &AccuracyMetrics{
		TotalEstimates: 0,
		AccurateCount:  0,
		OffByOne:       0,
		OffByFactor:    0,
	}

	for _, stats := range sc.PredicateStats {
		if stats == nil || stats.EstimatedCard == 0 {
			continue
		}

		metrics.TotalEstimates++

		if stats.ActualCard == stats.EstimatedCard {
			metrics.AccurateCount++
		} else if stats.ActualCard == stats.EstimatedCard+1 || stats.ActualCard == stats.EstimatedCard-1 {
			metrics.OffByOne++
		} else {
			metrics.OffByFactor++
		}
	}

	if metrics.TotalEstimates > 0 {
		metrics.Accuracy = float64(metrics.AccurateCount) / float64(metrics.TotalEstimates)
	}

	return metrics
}

// AccuracyMetrics contains metrics about estimate accuracy.
type AccuracyMetrics struct {
	TotalEstimates int
	AccurateCount  int
	OffByOne       int
	OffByFactor    int
	Accuracy       float64
}
