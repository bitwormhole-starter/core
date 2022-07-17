package collection

// Arguments ...
type Arguments interface {
	Count() int

	GetArgument(index int) string

	Export(dst []string) []string

	Import(src []string)
}
