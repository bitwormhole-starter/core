package collection

// Resource ...
type Resource interface {
}

// Resources ...
type Resources interface {
	GetResource(name string) Resource

	SetResource(name string, value Resource)

	Export(dst map[string]Resource) map[string]Resource

	Import(src map[string]Resource)
}
