package collection

// Properties ...
type Properties interface {
	GetProperty(name string) string

	SetProperty(name string, value string)

	Export(dst map[string]string) map[string]string

	Import(src map[string]string)
}
