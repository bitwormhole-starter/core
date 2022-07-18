package collection

import (
	"strings"

	"bitwormhole.com/starter/core/generic"
)

// Resource ...
type Resource interface {
}

// Resources ...
type Resources interface {
	GetResource(name string) Resource

	SetResource(name string, value Resource)

	Export(dst map[string]Resource) map[string]Resource

	Import(src map[string]Resource, reset bool)
}

////////////////////////////////////////////////////////////////////////////////

// DefaultResources ...
type DefaultResources struct {
	table generic.Map[string, Resource]
}

func (inst *DefaultResources) _Impl() Resources {
	return inst
}

func (inst *DefaultResources) getTab() generic.Map[string, Resource] {
	table := inst.table
	if table == nil {
		table = generic.NewMap[string, Resource]()
		inst.table = table
	}
	return table
}

func (inst *DefaultResources) normalizeResourceName(name string) string {
	const sep = "/"
	const prefix = ":" + sep
	name = strings.ReplaceAll(name, "\\", sep)
	// remove prefix
	index := strings.Index(name, prefix)
	if index >= 0 {
		name = name[index+len(prefix):]
	}
	// to array
	elements := strings.Split(name, sep)
	// build
	builder := strings.Builder{}
	for _, el := range elements {
		el = strings.TrimSpace(el)
		if el == "" {
			continue
		}
		builder.WriteString(sep)
		builder.WriteString(el)
	}
	return builder.String()
}

// GetResource ...
func (inst *DefaultResources) GetResource(name string) Resource {
	name = inst.normalizeResourceName(name)
	return inst.getTab().Get(name)
}

// SetResource ...
func (inst *DefaultResources) SetResource(name string, value Resource) {
	name = inst.normalizeResourceName(name)
	inst.getTab().Put(name, value)
}

// Export ...
func (inst *DefaultResources) Export(dst map[string]Resource) map[string]Resource {
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
func (inst *DefaultResources) Import(src map[string]Resource, reset bool) {
	mid := make(map[string]Resource, 0)
	for k, v := range src {
		k = inst.normalizeResourceName(k)
		mid[k] = v
	}
	table := inst.getTab()
	if reset {
		table.Clear()
	}
	table.PutAll(mid)
}
