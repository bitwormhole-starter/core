package generic

import "errors"

////////////////////////////////////////////////////////////////////////////////
// ArrayList

// ArrayList 用数组实现一个 List 接口
type ArrayList[T any] struct {
	array      []T
	null       T
	nilChecker NilChecker[T]
}

func (inst *ArrayList[T]) _Impl() List[T] {
	return inst
}

func (inst *ArrayList[T]) Length() int {
	list := inst.array
	if list == nil {
		return 0
	}
	return len(list)
}

func (inst *ArrayList[T]) Add(item T) {
	inst.array = append(inst.array, item)
}

func (inst *ArrayList[T]) AddAll(src []T) {
	inst.array = append(inst.array, src...)
}

func (inst *ArrayList[T]) ToArray() []T {
	src := inst.array
	dst := make([]T, 0)
	for _, item := range src {
		dst = append(dst, item)
	}
	return dst
}

func (inst *ArrayList[T]) Clear() {
	inst.array = make([]T, 0)
}

func (inst *ArrayList[T]) GetRequired(index int) (T, error) {
	list := inst.array
	if list == nil {
		return inst.null, errors.New("list is nil")
	}
	size := len(list)
	if 0 <= index && index < size {
		item := list[index]
		if inst.nilChecker.IsNil(item) {
			return inst.null, errors.New("item is nil")
		}
		return item, nil
	}
	return inst.null, errors.New("out of bounds")
}

func (inst *ArrayList[T]) GetOptional(index int, def T) T {
	list := inst.array
	if list == nil {
		return inst.null
	}
	size := len(list)
	if 0 <= index && index < size {
		item := list[index]
		if inst.nilChecker.IsNil(item) {
			item = def
		}
		return item
	}
	return def
}

func (inst *ArrayList[T]) Get(index int) T {
	return inst.GetOptional(index, inst.null)
}

////////////////////////////////////////////////////////////////////////////////
