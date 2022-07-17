package collection

// Environment ...
type Environment interface {
	GetEnv(name string) string

	SetEnv(name string, value string)

	Export(dst map[string]string) map[string]string

	Import(src map[string]string)
}
