package collection

// Parameters ...
type Parameters interface {
	GetParameter(name string) string

	SetParameter(name string, value string)

	Export(dst map[string]string) map[string]string

	Import(src map[string]string)
}
