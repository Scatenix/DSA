package linkedList

import "errors"

var ErrIndexOutOfRange = errors.New("index out of range")
var ErrAccessEmptyList = errors.New("tried to remove from an empty list")

type Node[T any] struct {
	Value any
	Next  *Node[T]
}

type LinkedList[T any] struct {
	Head *Node[T]
	Tail *Node[T]
	Size uint
}

// NewLinkedList creates a new linkedList with n elements of type T.
func NewLinkedList[T any](elems ...T) *LinkedList[T] {

	ll := LinkedList[T]{Head: nil, Tail: nil, Size: 0}

	for _, e := range elems {
		ll.Push(e)
	}
	return &ll
}

// IsEmpty returns true if the linkedList does not contain any values.
func (l *LinkedList[T]) IsEmpty() bool {
	if l.Size <= 0 {
		return true
	}
	return false
}

// Get the value at i.
//
// returns value at i.
//
// returns error if i is out of bounds.
func (l *LinkedList[T]) Get(i uint) (val T, err error) {
	node, err := l.GetNode(i)

	if err != nil {
		var ret T
		return ret, errors.New(err.Error())
	}
	return node.Value.(T), nil
}

// GetNode at i.
//
// Returns the Node at i.
//
// Returns error if i is out of bounds.
func (l *LinkedList[T]) GetNode(i uint) (node *Node[T], err error) {
	if i == l.Size-1 {
		return l.Tail, nil
	}
	if i >= l.Size {
		return nil, ErrIndexOutOfRange
	}
	curNode := l.Head
	for j := uint(0); j < i; j++ {
		curNode = curNode.Next
	}
	return curNode, nil
}

// Pop removes the head of the list.
//
// Returns its value.
//
// Returns error if there was no head.
func (l *LinkedList[T]) Pop() (val T, err error) {
	var ret T
	if l.Head == nil {
		return ret, ErrIndexOutOfRange
	}

	oldHead := l.Head
	ret = oldHead.Value.(T)
	l.Head = l.Head.Next
	oldHead = nil
	l.Size--

	if l.Head == nil {
		l.Tail = nil
	}

	return ret, nil
}

// Push inserts the value at the end of the list as new tail.
//
// Returns the pushed Node.
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

// PushFront inserts the value at the front of the list as new head.
//
// Returns the pushed Node.
func (l *LinkedList[T]) PushFront(val T) *Node[T] {
	oldHead := l.Head
	newHead := Node[T]{val, oldHead}
	l.Head = &newHead
	l.Size++

	if l.Tail == nil {
		l.Tail = &newHead
	}

	return &newHead
}

// Insert the value at i.
//
// Returns the new inserted Node.
// Returns error if i is out of range.
func (l *LinkedList[T]) Insert(val T, i uint) (node *Node[T], err error) {
	if i > l.Size {
		return nil, ErrIndexOutOfRange
	}

	if l.IsEmpty() {
		newNode := Node[T]{val, nil}
		l.Head = &newNode
		l.Tail = &newNode
		l.Size++
		return &newNode, nil
	}
	l.Size++
	if i == 0 {
		nextNode, _ := l.GetNode(0)
		newNode := Node[T]{val, nextNode}
		l.Head = &newNode
		return &newNode, nil
	}

	prevNode, _ := l.GetNode(i - 1)
	nextNode := prevNode.Next
	newNode := Node[T]{val, nextNode}
	prevNode.Next = &newNode

	if i == l.Size-1 {
		l.Tail = &newNode
	}

	return &newNode, nil
}

// Remove the Node at i.
//
// Returns the removed value at i.
//
// Returns an error if list is empty or i >= l.Size.
func (l *LinkedList[T]) Remove(i uint) (val T, err error) {
	var value T
	if l.IsEmpty() {
		return value, ErrAccessEmptyList
	} else if i >= l.Size {
		return value, ErrIndexOutOfRange
	}

	if i == 0 {
		value = l.Head.Value.(T)
		newHead := l.Head.Next
		l.Head = nil
		l.Head = newHead
		l.Size--
		if l.Head == nil {
			l.Tail = nil
		}
		return value, err
	}
	prevNode, _ := l.GetNode(i - 1)
	oldNode := prevNode.Next
	newNextNode := oldNode.Next

	value = prevNode.Next.Value.(T)

	oldNode = nil
	prevNode.Next = newNextNode
	if i == l.Size-1 {
		l.Tail = prevNode
	}

	l.Size--
	return value, nil
}
