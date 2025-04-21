package dto

type Sort struct {
	ColId string `json:"col_id"`
	Sort  string `json:"sort"`
}

type Filter struct {
	Type       string `json:"type"`
	From       string `json:"from"`
	To         string `json:"to"`
	FilterType string `json:"filter_type"`
}

type DynamicFilter struct {
	Sort   *[]Sort           `json:"sort"`
	Filter map[string]Filter `json:"filter"`
}

type PageList[T any] struct {
	PageNumber      int   `json:"page_number"`
	TotalRows       int64 `json:"total_rows"`
	TotalPages      int   `json:"total_pages"`
	HasPreviousPage bool  `json:"has_previous_page"`
	HasNextPage     bool  `json:"has_next_page"`
	Items           *[]T  `json:"items"`
}

type PaginationInput struct {
	PageSize   int `json:"page_size"`
	PageNumber int `json:"page_number"`
}

type PaginationInputWithFilter struct {
	PaginationInput
	DynamicFilter
}

func (p *PaginationInputWithFilter) GetOffset() int {
	return (p.PageNumber - 1) * p.PageSize
}

func (p *PaginationInputWithFilter) GetPageSize() int {
	if p.PageSize == 0 {
		p.PageSize = 10
	}

	return p.PageSize
}

func (p *PaginationInputWithFilter) GetPageNumber() int {
	if p.PageNumber == 0 {
		p.PageNumber = 1
	}

	return p.PageNumber
}
