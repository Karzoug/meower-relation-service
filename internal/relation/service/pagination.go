package service

type pagination struct {
	pageToken   string
	maxPageSize int
}

func defaultPagination() pagination {
	return pagination{
		pageToken:   "",
		maxPageSize: 20,
	}
}

type PaginationOption func(*pagination)

func WithPaginationToken(token string) PaginationOption {
	return func(p *pagination) {
		p.pageToken = token
	}
}

func WithPaginationMaxPageSize(size int) PaginationOption {
	if size <= 0 {
		size = 20
	}
	if size > 100 {
		size = 100
	}
	return func(p *pagination) {
		p.maxPageSize = size
	}
}

// WithoutPagination is a pagination option that disables pagination,
// returning all results.
func WithoutPagination() PaginationOption {
	return func(p *pagination) {
		p.pageToken = ""
		p.maxPageSize = -1
	}
}
