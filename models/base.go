package models

type BaseResponse[T any] struct {
	Message string          `json:"message"`
	Data    T               `json:"data,omitempty"`
	Errors  []any           `json:"errors,omitempty"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
}
