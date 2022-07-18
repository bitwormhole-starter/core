package generic

// NewMap 创建一个简单的字典
func NewMap[K int | string, V any]() Map[K, V] {
	return &GoMap[K, V]{}
}

// NewList 创建一个简单的列表
func NewList[T any]() List[T] {
	return &ArrayList[T]{}
}

// NewQueue 创建一个简单的队列
func NewQueue[T any]() Queue[T] {
	panic("no impl")
	// return nil
}

// NewStack 创建一个简单的堆栈
func NewStack[T any]() Stack[T] {
	panic("no impl")
	// return nil
}
