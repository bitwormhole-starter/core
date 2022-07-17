package collection

// Attributes ...
type Attributes interface {
	GetAttribute(name string) any

	SetAttribute(name string, value any)

	Export(dst map[string]any) map[string]any

	Import(src map[string]any)
}
