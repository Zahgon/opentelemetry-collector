package component

type Host interface {
	GetExtensions() map[ID]Component
}
