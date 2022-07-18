package generic

// NilChecker 用来检查一个数值是否为空
type NilChecker[T any] interface {
	IsNil(o T) bool
}

// List 表示一个抽象的列表
type List[T any] interface {
	Add(o T)

	AddAll(src []T)

	Clear()

	Length() int

	Get(index int) T

	GetRequired(index int) (T, error)

	GetOptional(index int, def T) T

	ToArray() []T
}

// Map 表示一个抽象的字典
type Map[K int | string, V any] interface {
	Put(key K, value V)

	PutAll(src map[K]V)

	Get(key K) V

	Clear()

	ToMap() map[K]V
}

// Stack 表示一个抽象的堆栈
type Stack[T any] interface {
	Clear()

	Push(o T)

	Pop() T
}

// Queue 表示一个抽象的队列
type Queue[T any] interface {
	Clear()

	Push(o T)

	PushAll(src []T)

	Pop() T

	ToArray() []T
}
