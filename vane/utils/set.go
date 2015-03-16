package utils

type StringSet map[string]struct{}

func NewStringSet(items ...string) *StringSet {
	set := make(StringSet, len(items))
	for _, item := range items {
		set[item] = struct{}{}
	}
	return &set
}

func (set *StringSet) Contains(item string) bool {
	_, contains := (*set)[item]
	return contains
}
