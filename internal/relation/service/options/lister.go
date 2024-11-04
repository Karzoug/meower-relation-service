package options

// Lister is an interface that wraps a List method to return a
// slice of option setters.
type Lister[T any] interface {
	List() []func(*T)
}

// GetOptions will functionally merge a slice of option setters in a
// "last-one-wins" manner, where nil option setters are ignored.
func GetOptions[T any](opts ...Lister[T]) (*T, error) {
	args := new(T)
	for _, opt := range opts {
		if opt == nil {
			continue
		}

		for _, setArgs := range opt.List() {
			if setArgs == nil {
				continue
			}
			setArgs(args)
		}
	}
	return args, nil
}
