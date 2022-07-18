package collection

// Properties ...
type Properties interface {
	GetProperty(name string) string

	SetProperty(name string, value string)

	Export(dst map[string]string) map[string]string

	Import(src map[string]string, reset bool)
}

////////////////////////////////////////////////////////////////////////////////

// DefaultProperties ...
type DefaultProperties struct {
	table mapStrStr
}

func (inst *DefaultProperties) _Impl() Properties {
	return inst
}

// GetProperty ...
func (inst *DefaultProperties) GetProperty(name string) string {
	return inst.table.get(name)
}

// SetProperty ...
func (inst *DefaultProperties) SetProperty(name string, value string) {
	inst.table.put(name, value)
}

// Export ...
func (inst *DefaultProperties) Export(dst map[string]string) map[string]string {
	return inst.table.doExport(dst)
}

// Import ...
func (inst *DefaultProperties) Import(src map[string]string, reset bool) {
	inst.table.doImport(src, reset)
}
