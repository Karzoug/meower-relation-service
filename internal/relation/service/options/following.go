package options

type ListFollowingsOptions struct {
	PageToken   string
	MaxPageSize int
}

type listFollowingsOptionsBuilder struct {
	Opts []func(*ListFollowingsOptions)
}

func ListFollowings() *listFollowingsOptionsBuilder {
	return &listFollowingsOptionsBuilder{}
}

func (f *listFollowingsOptionsBuilder) List() []func(*ListFollowingsOptions) {
	return f.Opts
}

func (f *listFollowingsOptionsBuilder) SetPagination(pageToken string, maxPageSize int) *listFollowingsOptionsBuilder {
	f.Opts = append(f.Opts, func(opts *ListFollowingsOptions) {
		if maxPageSize <= 0 {
			maxPageSize = 20
		}
		if maxPageSize > 100 {
			maxPageSize = 100
		}

		opts.PageToken = pageToken
		opts.MaxPageSize = maxPageSize
	})
	return f
}

func (f *listFollowingsOptionsBuilder) DisablePagination() *listFollowingsOptionsBuilder {
	f.Opts = append(f.Opts, func(opts *ListFollowingsOptions) {
		opts.PageToken = ""
		opts.MaxPageSize = -1
	})
	return f
}
