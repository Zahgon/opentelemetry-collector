package filter

type Filter interface {
	Matches(any) bool
}
