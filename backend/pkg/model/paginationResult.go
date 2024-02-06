package model

type PaginationResult struct {
	TotalCount int64       `json:"total_count"`
	TotalPages int32       `json:"total_pages"`
	Page       int32       `json:"page"`
	PageSize   int32       `json:"page_size"`
	Data       interface{} `json:"data"`
}
