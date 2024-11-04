package options

type ListFollowersOptions struct {
	IncludeHidden bool
	PageToken     string
	MaxPageSize   int
}

type listFollowersOptionsBuilder struct {
	Opts []func(*ListFollowersOptions)
}

func ListFollowers() *listFollowersOptionsBuilder {
	return &listFollowersOptionsBuilder{}
}

func (f *listFollowersOptionsBuilder) List() []func(*ListFollowersOptions) {
	return f.Opts
}

func (f *listFollowersOptionsBuilder) SetPagination(pageToken string, maxPageSize int) *listFollowersOptionsBuilder {
	f.Opts = append(f.Opts, func(opts *ListFollowersOptions) {
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

func (f *listFollowersOptionsBuilder) DisablePagination() *listFollowersOptionsBuilder {
	f.Opts = append(f.Opts, func(opts *ListFollowersOptions) {
		opts.PageToken = ""
		opts.MaxPageSize = -1
	})
	return f
}

func (f *listFollowersOptionsBuilder) IncludeHidden() *listFollowersOptionsBuilder {
	f.Opts = append(f.Opts, func(opts *ListFollowersOptions) {
		opts.IncludeHidden = true
	})
	return f
}
