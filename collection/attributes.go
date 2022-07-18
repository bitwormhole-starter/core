package collection

import "bitwormhole.com/starter/core/generic"

// Attributes ...
type Attributes interface {
	GetAttribute(name string) any

	SetAttribute(name string, value any)

	Export(dst map[string]any) map[string]any

	Import(src map[string]any, reset bool)
}

////////////////////////////////////////////////////////////////////////////////

// DefaultAttributes ...
type DefaultAttributes struct {
	table generic.Map[string, any]
}

func (inst *DefaultAttributes) _Impl() Attributes {
	return inst
}

func (inst *DefaultAttributes) getTab() generic.Map[string, any] {
	table := inst.table
	if table == nil {
		table = generic.NewMap[string, any]()
		inst.table = table
	}
	return table
}

// GetAttribute ...
func (inst *DefaultAttributes) GetAttribute(name string) any {
	return inst.getTab().Get(name)
}

// SetAttribute ...
func (inst *DefaultAttributes) SetAttribute(name string, value any) {
	inst.getTab().Put(name, value)
}

// Export ...
func (inst *DefaultAttributes) Export(dst map[string]any) map[string]any {
	src := inst.getTab().ToMap()
	if dst == nil {
		return src
	}
	for k, v := range src {
		dst[k] = v
	}
	return dst
}

// Import ...
func (inst *DefaultAttributes) Import(src map[string]any, reset bool) {
	table := inst.getTab()
	if reset {
		table.Clear()
	}
	table.PutAll(src)
}
