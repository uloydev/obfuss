package models

type TodoRequest struct {
	Action string `json:"action" binding:"required"`
}
