package models

type PaginationMeta struct {
	Total      int `json:"total"`
	PageNumber int `json:"page_number"`
	PageSize   int `json:"page_size"`
}
