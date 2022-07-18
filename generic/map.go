package generic

// GoMap 用 go-map[k]v 实现一个简单的字典
type GoMap[K int | string, V any] struct {
	table map[K]V
}

func (inst *GoMap[K, V]) _Impl() Map[K, V] {
	return inst
}

func (inst *GoMap[K, V]) getTab() map[K]V {
	tab := inst.table
	if tab == nil {
		tab = make(map[K]V)
		inst.table = tab
	}
	return tab
}

func (inst *GoMap[K, V]) Put(key K, value V) {
	tab := inst.getTab()
	tab[key] = value
}

func (inst *GoMap[K, V]) PutAll(src map[K]V) {
	dst := inst.getTab()
	for k, v := range src {
		dst[k] = v
	}
}

func (inst *GoMap[K, V]) Get(key K) V {
	tab := inst.getTab()
	return tab[key]
}

func (inst *GoMap[K, V]) Clear() {
	inst.table = nil
}

func (inst *GoMap[K, V]) ToMap() map[K]V {
	src := inst.getTab()
	dst := make(map[K]V)
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
