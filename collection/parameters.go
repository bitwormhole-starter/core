package collection

// Parameters ...
type Parameters interface {
	GetParameter(name string) string

	SetParameter(name string, value string)

	Export(dst map[string]string) map[string]string

	Import(src map[string]string, reset bool)
}

////////////////////////////////////////////////////////////////////////////////

// DefaultParameters ...
type DefaultParameters struct {
	table mapStrStr
}

func (inst *DefaultParameters) _Impl() Parameters {
	return inst
}

// GetParameter ...
func (inst *DefaultParameters) GetParameter(name string) string {
	return inst.table.get(name)
}

// SetParameter ...
func (inst *DefaultParameters) SetParameter(name string, value string) {
	inst.table.put(name, value)
}

// Export ...
func (inst *DefaultParameters) Export(dst map[string]string) map[string]string {
	return inst.table.doExport(dst)
}

// Import ...
func (inst *DefaultParameters) Import(src map[string]string, reset bool) {
	inst.table.doImport(src, reset)
}
