package sparql

import (
	"testing"
)

func TestStatisticsCollectorRecordExecution(t *testing.T) {
	sc := NewStatisticsCollector()

	sc.RecordExecution("name", 100)
	sc.RecordExecution("name", 100)
	sc.RecordExecution("age", 50)

	stats := sc.GetStats("name")
	if stats.ExecutionCount != 2 {
		t.Errorf("expected 2 executions for 'name', got %d", stats.ExecutionCount)
	}

	if stats.ActualCard != 100 {
		t.Errorf("expected cardinality 100, got %d", stats.ActualCard)
	}

	if sc.ExecutionCount != 3 {
		t.Errorf("expected 3 total executions, got %d", sc.ExecutionCount)
	}
}

func TestStatisticsCollectorRecordIndexUsage(t *testing.T) {
	sc := NewStatisticsCollector()

	sc.RecordIndexUsage("name", "hash")
	sc.RecordIndexUsage("name", "hash")
	sc.RecordIndexUsage("age", "range")

	stats := sc.GetStats("name")
	if stats.IndexUsageCount != 2 {
		t.Errorf("expected 2 index usages for 'name', got %d", stats.IndexUsageCount)
	}

	idxStats := sc.IndexStats["name"]
	if idxStats == nil {
		t.Errorf("expected index stats for 'name'")
	}

	if idxStats.UsageCount != 2 {
		t.Errorf("expected 2 usages, got %d", idxStats.UsageCount)
	}
}

func TestStatisticsCollectorRecordJoinExecution(t *testing.T) {
	sc := NewStatisticsCollector()

	sc.RecordJoinExecution("name", "age", []string{"?x"}, 0.3)
	sc.RecordJoinExecution("name", "age", []string{"?x"}, 0.3)
	sc.RecordJoinExecution("email", "phone", []string{"?x"}, 0.2)

	if len(sc.JoinStats) != 2 {
		t.Errorf("expected 2 distinct joins, got %d", len(sc.JoinStats))
	}

	// Check that same join is counted
	if sc.JoinStats[0].ExecutionCount != 2 {
		t.Errorf("expected 2 executions for first join, got %d", sc.JoinStats[0].ExecutionCount)
	}
}

func TestStatisticsCollectorGetStats(t *testing.T) {
	sc := NewStatisticsCollector()

	sc.RecordExecution("name", 100)

	stats := sc.GetStats("name")
	if stats.Predicate != "name" {
		t.Errorf("expected predicate 'name', got %s", stats.Predicate)
	}

	// Get non-existent predicate
	stats = sc.GetStats("unknown")
	if stats.ExecutionCount != 0 {
		t.Errorf("expected 0 executions for unknown predicate")
	}
}

func TestStatisticsCollectorTypeStats(t *testing.T) {
	sc := NewStatisticsCollector()

	sc.RecordTypeInstance("Person", 1000)
	sc.RecordTypeInstance("Person", 2000)

	stats := sc.GetTypeStats("Person")
	if stats.InstanceCount != 2000 {
		t.Errorf("expected 2000 instances, got %d", stats.InstanceCount)
	}
}

func TestStatisticsCollectorUpdateEstimates(t *testing.T) {
	sc := NewStatisticsCollector()

	sc.RecordExecution("name", 100)
	sc.RecordExecution("name", 100)

	sc.UpdateEstimates()

	stats := sc.GetStats("name")
	// After update, selectivity should be calculated
	if stats.SelectivityActual == 0 {
		t.Errorf("expected selectivity actual to be calculated")
	}

	if stats.SelectivityEstimate == 0 {
		t.Errorf("expected selectivity estimate to be set from actual")
	}

	if stats.SelectivityEstimate != stats.SelectivityActual {
		t.Errorf("expected estimate to match actual initially")
	}
}

func TestStatisticsCollectorGetStatisticsInfo(t *testing.T) {
	sc := NewStatisticsCollector()

	sc.RecordExecution("name", 100)
	sc.RecordExecution("age", 50)
	sc.RecordTypeInstance("Person", 1000)
	sc.RecordIndexUsage("name", "hash")
	sc.RecordJoinExecution("name", "age", []string{"?x"}, 0.3)

	info := sc.GetStatisticsInfo()

	if info.TrackedPredicates != 2 {
		t.Errorf("expected 2 tracked predicates, got %d", info.TrackedPredicates)
	}

	if info.TrackedTypes != 1 {
		t.Errorf("expected 1 tracked type, got %d", info.TrackedTypes)
	}

	if info.TrackedIndexes != 1 {
		t.Errorf("expected 1 tracked index, got %d", info.TrackedIndexes)
	}

	if info.TrackedJoins != 1 {
		t.Errorf("expected 1 tracked join, got %d", info.TrackedJoins)
	}

	if info.TotalExecutions != 2 {
		t.Errorf("expected 2 total executions, got %d", info.TotalExecutions)
	}
}

func TestStatisticsCollectorReset(t *testing.T) {
	sc := NewStatisticsCollector()

	sc.RecordExecution("name", 100)
	sc.RecordExecution("age", 50)

	info := sc.GetStatisticsInfo()
	if info.TrackedPredicates != 2 {
		t.Errorf("expected 2 tracked predicates before reset")
	}

	sc.Reset()

	info = sc.GetStatisticsInfo()
	if info.TrackedPredicates != 0 {
		t.Errorf("expected 0 tracked predicates after reset")
	}

	if sc.ExecutionCount != 0 {
		t.Errorf("expected 0 total executions after reset")
	}
}

func TestStatisticsCollectorGetAccuracyMetrics(t *testing.T) {
	sc := NewStatisticsCollector()

	// Create a stat with estimate and actual
	stat := &PredicateExecutionStats{
		Predicate:      "name",
		ExecutionCount: 1,
		EstimatedCard:  100,
		ActualCard:     100,
	}
	sc.PredicateStats["name"] = stat

	metrics := sc.GetAccuracyMetrics()

	if metrics.TotalEstimates != 1 {
		t.Errorf("expected 1 estimate, got %d", metrics.TotalEstimates)
	}

	if metrics.AccurateCount != 1 {
		t.Errorf("expected 1 accurate estimate, got %d", metrics.AccurateCount)
	}

	if metrics.Accuracy != 1.0 {
		t.Errorf("expected 100%% accuracy, got %f", metrics.Accuracy)
	}
}

func TestStatisticsCollectorGetAccuracyMetricsInaccurate(t *testing.T) {
	sc := NewStatisticsCollector()

	stat1 := &PredicateExecutionStats{
		Predicate:     "name",
		EstimatedCard: 100,
		ActualCard:    100,
	}
	sc.PredicateStats["name"] = stat1

	stat2 := &PredicateExecutionStats{
		Predicate:     "age",
		EstimatedCard: 50,
		ActualCard:    100,
	}
	sc.PredicateStats["age"] = stat2

	metrics := sc.GetAccuracyMetrics()

	if metrics.TotalEstimates != 2 {
		t.Errorf("expected 2 estimates")
	}

	if metrics.AccurateCount != 1 {
		t.Errorf("expected 1 accurate")
	}

	if metrics.Accuracy != 0.5 {
		t.Errorf("expected 50%% accuracy")
	}
}

func TestStatisticsCollectorEmptyStats(t *testing.T) {
	sc := NewStatisticsCollector()

	stats := sc.GetStats("name")
	if stats.ExecutionCount != 0 {
		t.Errorf("expected 0 executions for non-existent")
	}

	typeStats := sc.GetTypeStats("Person")
	if typeStats.InstanceCount != 0 {
		t.Errorf("expected 0 instances for non-existent type")
	}

	info := sc.GetStatisticsInfo()
	if info.TrackedPredicates != 0 {
		t.Errorf("expected 0 tracked predicates")
	}
}

func TestStatisticsCollectorMultipleIndexTypes(t *testing.T) {
	sc := NewStatisticsCollector()

	sc.RecordIndexUsage("name", "hash")
	sc.RecordIndexUsage("age", "range")
	sc.RecordIndexUsage("bio", "fulltext")

	if len(sc.IndexStats) != 3 {
		t.Errorf("expected 3 index stats, got %d", len(sc.IndexStats))
	}

	if sc.IndexStats["name"].IndexType != "hash" {
		t.Errorf("expected 'hash' index for name")
	}

	if sc.IndexStats["age"].IndexType != "range" {
		t.Errorf("expected 'range' index for age")
	}
}

func TestStatisticsCollectorSelectivityTracking(t *testing.T) {
	sc := NewStatisticsCollector()

	sc.RecordExecution("name", 100)
	stats := sc.GetStats("name")

	if stats.ActualCard != 100 {
		t.Errorf("expected actual cardinality 100, got %d", stats.ActualCard)
	}

	// After update, selectivity should be calculated
	sc.UpdateEstimates()

	stats = sc.GetStats("name")
	if stats.SelectivityEstimate == 0 && stats.SelectivityActual > 0 {
		t.Errorf("expected estimate to be updated from actual")
	}
}

func TestStatisticsCollectorJoinVariables(t *testing.T) {
	sc := NewStatisticsCollector()

	vars := []string{"?x", "?y"}
	sc.RecordJoinExecution("pred1", "pred2", vars, 0.5)

	if len(sc.JoinStats) != 1 {
		t.Errorf("expected 1 join")
	}

	joinStat := sc.JoinStats[0]
	if len(joinStat.JoinVars) != 2 {
		t.Errorf("expected 2 join variables, got %d", len(joinStat.JoinVars))
	}

	if joinStat.JoinVars[0] != "?x" || joinStat.JoinVars[1] != "?y" {
		t.Errorf("expected ['?x', '?y'], got %v", joinStat.JoinVars)
	}
}

func TestStatisticsCollectorEmptyRecords(t *testing.T) {
	sc := NewStatisticsCollector()

	// Test with empty strings
	sc.RecordExecution("", 100)
	sc.RecordIndexUsage("", "hash")

	info := sc.GetStatisticsInfo()
	if info.TrackedPredicates != 0 {
		t.Errorf("expected no stats for empty predicate")
	}
}
