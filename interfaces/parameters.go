package interfaces

type MongoFilter struct {
	Pagination
	Sort
	Projection any
}

type Pagination struct {
	PerPage int `json:"per_page"`
	Page    int `json:"page"`
}
type Sort struct {
	SortField string `json:"sort_field"`
	Asc       bool   `json:"asc"`
}
