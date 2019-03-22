package model

// Paging Model
type Paging struct {
	Skip      *int64
	Limit     *int64
	SortKey   string
	SortVal   int
	Condition interface{}
}
