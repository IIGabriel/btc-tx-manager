package interfaces

type MongoFilter struct {
	Pagination
	Sort
	Projection any
}

type Pagination struct {
	PerPage int `json:"perPage"`
	Page    int `json:"page"`
}
type Sort struct {
	Field string `json:"field"`
	Asc   bool   `json:"asc"`
}
