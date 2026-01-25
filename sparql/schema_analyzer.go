package sparql

// PredicateInfo contains metadata about a SPARQL predicate.
type PredicateInfo struct {
	Name          string
	Type          string
	ReverseIndex  bool
	IndexType     string
	Cardinality   uint64
	ListType      bool
	Lang          bool
	Upsert        bool
	Noconflict    bool
	Nullable      bool
	Count         bool
	PredicateType string
}

// TypeInfo contains metadata about a SPARQL type.
type TypeInfo struct {
	Name       string
	Fields     map[string]*PredicateInfo
	Predicates []string
}

// SchemaInfo contains the parsed Dgraph schema.
type SchemaInfo struct {
	Predicates map[string]*PredicateInfo
	Types      map[string]*TypeInfo
}

// SchemaAnalyzer analyzes Dgraph schema for query optimization.
type SchemaAnalyzer struct {
	Schema *SchemaInfo
}

// NewSchemaAnalyzer creates a new schema analyzer.
func NewSchemaAnalyzer(schema *SchemaInfo) *SchemaAnalyzer {
	if schema == nil {
		schema = &SchemaInfo{
			Predicates: make(map[string]*PredicateInfo),
			Types:      make(map[string]*TypeInfo),
		}
	}
	return &SchemaAnalyzer{
		Schema: schema,
	}
}

// GetPredicate returns metadata about a predicate.
func (sa *SchemaAnalyzer) GetPredicate(name string) *PredicateInfo {
	if sa == nil || sa.Schema == nil {
		return nil
	}
	return sa.Schema.Predicates[name]
}

// HasPredicate checks if a predicate exists in the schema.
func (sa *SchemaAnalyzer) HasPredicate(name string) bool {
	if sa == nil || sa.Schema == nil {
		return false
	}
	_, exists := sa.Schema.Predicates[name]
	return exists
}

// IsIndexed checks if a predicate has a specific index type.
func (sa *SchemaAnalyzer) IsIndexed(pred, indexType string) bool {
	if sa == nil || sa.Schema == nil {
		return false
	}

	predInfo := sa.Schema.Predicates[pred]
	if predInfo == nil {
		return false
	}

	if indexType == "" {
		return predInfo.IndexType != ""
	}

	return predInfo.IndexType == indexType
}

// GetIndexType returns the index type for a predicate.
func (sa *SchemaAnalyzer) GetIndexType(pred string) string {
	predInfo := sa.GetPredicate(pred)
	if predInfo == nil {
		return ""
	}
	return predInfo.IndexType
}

// GetCardinality returns the estimated cardinality for a predicate.
func (sa *SchemaAnalyzer) GetCardinality(pred string) uint64 {
	predInfo := sa.GetPredicate(pred)
	if predInfo == nil {
		return 0
	}
	return predInfo.Cardinality
}

// GetType returns metadata about a type.
func (sa *SchemaAnalyzer) GetType(name string) *TypeInfo {
	if sa == nil || sa.Schema == nil {
		return nil
	}
	return sa.Schema.Types[name]
}

// HasType checks if a type exists in the schema.
func (sa *SchemaAnalyzer) HasType(name string) bool {
	if sa == nil || sa.Schema == nil {
		return false
	}
	_, exists := sa.Schema.Types[name]
	return exists
}

// GetTypes returns all types in the schema.
func (sa *SchemaAnalyzer) GetTypes() []string {
	if sa == nil || sa.Schema == nil {
		return []string{}
	}

	types := make([]string, 0, len(sa.Schema.Types))
	for name := range sa.Schema.Types {
		types = append(types, name)
	}
	return types
}

// GetPredicates returns all predicates in the schema.
func (sa *SchemaAnalyzer) GetPredicates() []string {
	if sa == nil || sa.Schema == nil {
		return []string{}
	}

	predicates := make([]string, 0, len(sa.Schema.Predicates))
	for name := range sa.Schema.Predicates {
		predicates = append(predicates, name)
	}
	return predicates
}

// HasReverseIndex checks if a predicate has a reverse index.
func (sa *SchemaAnalyzer) HasReverseIndex(pred string) bool {
	predInfo := sa.GetPredicate(pred)
	if predInfo == nil {
		return false
	}
	return predInfo.ReverseIndex
}

// GetPredicateType returns the predicate type (scalar, uid, etc.).
func (sa *SchemaAnalyzer) GetPredicateType(pred string) string {
	predInfo := sa.GetPredicate(pred)
	if predInfo == nil {
		return ""
	}
	return predInfo.PredicateType
}

// IsListType checks if a predicate is a list type.
func (sa *SchemaAnalyzer) IsListType(pred string) bool {
	predInfo := sa.GetPredicate(pred)
	if predInfo == nil {
		return false
	}
	return predInfo.ListType
}

// IsLanguageSupported checks if a predicate supports language tags.
func (sa *SchemaAnalyzer) IsLanguageSupported(pred string) bool {
	predInfo := sa.GetPredicate(pred)
	if predInfo == nil {
		return false
	}
	return predInfo.Lang
}

// GetTypeFields returns all predicates associated with a type.
func (sa *SchemaAnalyzer) GetTypeFields(typeName string) []*PredicateInfo {
	typeInfo := sa.GetType(typeName)
	if typeInfo == nil {
		return []*PredicateInfo{}
	}

	fields := make([]*PredicateInfo, 0, len(typeInfo.Fields))
	for _, field := range typeInfo.Fields {
		fields = append(fields, field)
	}
	return fields
}

// HasCardinalityEstimate checks if cardinality is available for a predicate.
func (sa *SchemaAnalyzer) HasCardinalityEstimate(pred string) bool {
	predInfo := sa.GetPredicate(pred)
	if predInfo == nil {
		return false
	}
	return predInfo.Cardinality > 0
}

// GetSelectivity returns the selectivity of a predicate.
// Selectivity is a value between 0 and 1 indicating how selective the predicate is.
// Lower values = more selective (fewer matches per subject).
// Higher values = less selective (more matches per subject).
func (sa *SchemaAnalyzer) GetSelectivity(pred string) float64 {
	predInfo := sa.GetPredicate(pred)
	if predInfo == nil {
		// Default to medium selectivity
		return 0.5
	}

	// For indexed predicates, assume higher selectivity
	if predInfo.IndexType != "" {
		return 0.3
	}

	// For unindexed predicates, assume lower selectivity
	return 0.7
}

// IsCountSupported checks if count is supported for a predicate.
func (sa *SchemaAnalyzer) IsCountSupported(pred string) bool {
	predInfo := sa.GetPredicate(pred)
	if predInfo == nil {
		return false
	}
	return predInfo.Count
}

// AddPredicate adds a predicate to the schema.
func (sa *SchemaAnalyzer) AddPredicate(pred *PredicateInfo) {
	if sa == nil || sa.Schema == nil || pred == nil {
		return
	}
	sa.Schema.Predicates[pred.Name] = pred
}

// AddType adds a type to the schema.
func (sa *SchemaAnalyzer) AddType(typeInfo *TypeInfo) {
	if sa == nil || sa.Schema == nil || typeInfo == nil {
		return
	}
	sa.Schema.Types[typeInfo.Name] = typeInfo
}

// UpdateCardinality updates the cardinality for a predicate.
func (sa *SchemaAnalyzer) UpdateCardinality(pred string, cardinality uint64) {
	predInfo := sa.GetPredicate(pred)
	if predInfo == nil {
		return
	}
	predInfo.Cardinality = cardinality
}

// Copy creates a deep copy of the schema analyzer.
func (sa *SchemaAnalyzer) Copy() *SchemaAnalyzer {
	if sa == nil {
		return nil
	}

	newPredicates := make(map[string]*PredicateInfo)
	for name, pred := range sa.Schema.Predicates {
		if pred == nil {
			continue
		}
		newPred := *pred
		newPredicates[name] = &newPred
	}

	newTypes := make(map[string]*TypeInfo)
	for name, typeInfo := range sa.Schema.Types {
		if typeInfo == nil {
			continue
		}
		newTypeInfo := *typeInfo
		newTypeInfo.Fields = make(map[string]*PredicateInfo)
		for fieldName, field := range typeInfo.Fields {
			if field == nil {
				continue
			}
			newField := *field
			newTypeInfo.Fields[fieldName] = &newField
		}
		newTypeInfo.Predicates = make([]string, len(typeInfo.Predicates))
		copy(newTypeInfo.Predicates, typeInfo.Predicates)
		newTypes[name] = &newTypeInfo
	}

	return &SchemaAnalyzer{
		Schema: &SchemaInfo{
			Predicates: newPredicates,
			Types:      newTypes,
		},
	}
}

// PredicateCount returns the number of predicates in the schema.
func (sa *SchemaAnalyzer) PredicateCount() int {
	if sa == nil || sa.Schema == nil {
		return 0
	}
	return len(sa.Schema.Predicates)
}

// TypeCount returns the number of types in the schema.
func (sa *SchemaAnalyzer) TypeCount() int {
	if sa == nil || sa.Schema == nil {
		return 0
	}
	return len(sa.Schema.Types)
}

// IsEmpty checks if the schema is empty.
func (sa *SchemaAnalyzer) IsEmpty() bool {
	if sa == nil || sa.Schema == nil {
		return true
	}
	return len(sa.Schema.Predicates) == 0 && len(sa.Schema.Types) == 0
}

// GetIndexedPredicates returns all predicates that have an index.
func (sa *SchemaAnalyzer) GetIndexedPredicates() []string {
	if sa == nil || sa.Schema == nil {
		return []string{}
	}

	indexed := make([]string, 0)
	for name, pred := range sa.Schema.Predicates {
		if pred != nil && pred.IndexType != "" {
			indexed = append(indexed, name)
		}
	}
	return indexed
}

// GetUnindexedPredicates returns all predicates that don't have an index.
func (sa *SchemaAnalyzer) GetUnindexedPredicates() []string {
	if sa == nil || sa.Schema == nil {
		return []string{}
	}

	unindexed := make([]string, 0)
	for name, pred := range sa.Schema.Predicates {
		if pred != nil && pred.IndexType == "" {
			unindexed = append(unindexed, name)
		}
	}
	return unindexed
}

// GetPredicatesWithReverseIndex returns all predicates with reverse indexes.
func (sa *SchemaAnalyzer) GetPredicatesWithReverseIndex() []string {
	if sa == nil || sa.Schema == nil {
		return []string{}
	}

	reverse := make([]string, 0)
	for name, pred := range sa.Schema.Predicates {
		if pred != nil && pred.ReverseIndex {
			reverse = append(reverse, name)
		}
	}
	return reverse
}
