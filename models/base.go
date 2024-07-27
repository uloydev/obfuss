package models

type BaseResponse[T any] struct {
	Message string `json:"message" example:"success"`
	Data    T      `json:"data,omitempty"`
	Errors  []any  `json:"errors,omitempty"`
}
