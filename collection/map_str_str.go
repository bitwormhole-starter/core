package collection

import "bitwormhole.com/starter/core/generic"

type mapStrStr struct {
	table generic.Map[string, string]
}

func (inst *mapStrStr) getTab() generic.Map[string, string] {
	tab := inst.table
	if tab == nil {
		tab = generic.NewMap[string, string]()
		inst.table = tab
	}
	return tab
}

func (inst *mapStrStr) get(name string) string {
	return inst.getTab().Get(name)
}

func (inst *mapStrStr) put(name string, value string) {
	inst.getTab().Put(name, value)
}

func (inst *mapStrStr) doImport(src map[string]string, reset bool) {
	tab := inst.getTab()
	if reset {
		tab.Clear()
	}
	tab.PutAll(src)
}

func (inst *mapStrStr) doExport(dst map[string]string) map[string]string {
	if dst == nil {
		dst = make(map[string]string)
	}
	src := inst.getTab().ToMap()
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
