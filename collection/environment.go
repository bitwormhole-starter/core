package collection

// Environment ...
type Environment interface {
	GetEnv(name string) string

	SetEnv(name string, value string)

	Export(dst map[string]string) map[string]string

	Import(src map[string]string, reset bool)
}

////////////////////////////////////////////////////////////////////////////////

// DefaultEnvironment ...
type DefaultEnvironment struct {
	table mapStrStr
}

func (inst *DefaultEnvironment) _Impl() Environment {
	return inst
}

// GetEnv ...
func (inst *DefaultEnvironment) GetEnv(name string) string {
	return inst.table.get(name)
}

// SetEnv ...
func (inst *DefaultEnvironment) SetEnv(name string, value string) {
	inst.table.put(name, value)
}

// Export ...
func (inst *DefaultEnvironment) Export(dst map[string]string) map[string]string {
	return inst.table.doExport(dst)
}

// Import ...
func (inst *DefaultEnvironment) Import(src map[string]string, reset bool) {
	inst.table.doImport(src, reset)
}
