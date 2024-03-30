package models

type PaginationMeta struct {
	Total     int64 `json:"total"`
	TotalPage int64 `json:"totalPage"`
	Size      int   `json:"size"`
	Page      int   `json:"page"`
}

type PaginationParams struct {
	Page int `form:"page"`
	Size int `form:"size"`
}
