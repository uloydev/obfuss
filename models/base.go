package models

type BaseResponse[T any] struct {
	Message string          `json:"message"`
	Data    T               `json:"data"`
	Errors  []any           `json:"errors"`
	Meta    *PaginationMeta `json:"meta,omitempty"`
}
