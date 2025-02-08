package linkedList

type Node[T any] struct {
	Value any
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	size int
}

func NewLinkedList[T any](elems ...T) *LinkedList[T] {

	ll := LinkedList[T]{Head: nil, Tail: nil, size: 0}

	for _, e := range elems {
		ll.Push(e)
	}
	return &ll
}

func (l *LinkedList[T]) IsEmpty() bool {
	if l.size <= 0 {
		return true
	}
	return false
}

func (l *LinkedList[T]) Get(i int) (T, bool) {
	var value T
	return value, false
}

func (l *LinkedList[T]) Pop() (T, bool) {
	var value T
	return value, false
}

func (l *LinkedList[T]) Push(val T) {

}

func (l *LinkedList[T]) PushFront(val T) {

}

func (l *LinkedList[T]) Insert(val T, i int) {

}

func (l *LinkedList[T]) Remove(i int) {

}
