package collection

import "bitwormhole.com/starter/core/generic"

// Arguments ...
type Arguments interface {
	Count() int

	GetArgument(index int) string

	Export(dst []string) []string

	Import(src []string, reset bool)
}

////////////////////////////////////////////////////////////////////////////////

// DefaultArguments ...
type DefaultArguments struct {
	list generic.List[string]
}

func (inst *DefaultArguments) _Impl() Arguments {
	return inst
}

func (inst *DefaultArguments) getList() generic.List[string] {
	list := inst.list
	if list == nil {
		list = generic.NewList[string]()
		inst.list = list
	}
	return list
}

// Count ...
func (inst *DefaultArguments) Count() int {
	return inst.getList().Length()
}

// GetArgument ...
func (inst *DefaultArguments) GetArgument(index int) string {
	return inst.getList().Get(index)
}

// Export ...
func (inst *DefaultArguments) Export(dst []string) []string {
	src := inst.getList().ToArray()
	if dst == nil {
		return src
	}
	for _, item := range src {
		dst = append(dst, item)
	}
	return dst
}

// Import ...
func (inst *DefaultArguments) Import(src []string, reset bool) {
	list := inst.getList()
	if reset {
		list.Clear()
	}
	list.AddAll(src)
}
