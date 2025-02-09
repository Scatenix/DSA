package linkedList

type Node[T any] struct {
	Value any
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size int
}

func NewLinkedList[T any](elems ...T) *LinkedList[T] {

	ll := LinkedList[T]{Head: nil, Tail: nil, Size: 0}

	for _, e := range elems {
		ll.Push(e)
	}
	return &ll
}

func (l *LinkedList[T]) IsEmpty() bool {
	if l.Size <= 0 {
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

func (l *LinkedList[T]) Push(val T) *Node[T] {
	newNode := &Node[T]{val, nil}

	l.Size++

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return newNode
	}

	l.Tail.Next = newNode
	l.Tail = newNode

	return newNode
}

func (l *LinkedList[T]) PushFront(val T) {

}

func (l *LinkedList[T]) Insert(val T, i int) {

}

func (l *LinkedList[T]) Remove(i int) {

}
