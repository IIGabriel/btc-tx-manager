package models

import "encoding/json"

type HTTPResponse struct {
	Data    any    `json:"data"`
	PerPage uint64 `json:"perPage,omitempty"`
	Page    uint64 `json:"page,omitempty"`
	Total   uint64 `json:"total,omitempty"`
}

type HTTPErrorResponse struct {
	Error        json.RawMessage `json:"error,omitempty"`
	ErrorMessage string          `json:"errorMessage,omitempty"`
	Message      string          `json:"message"`
}
