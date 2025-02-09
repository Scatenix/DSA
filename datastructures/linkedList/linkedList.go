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
// returns false if the is no Node at i.
func (l *LinkedList[T]) Get(i int) (val T, success bool) {
	node, status := l.GetNode(i)

	if !status {
		var ret T
		return ret, false
	}
	return node.Value.(T), true
}

// GetNode at i.
//
// Returns the Node at i, nil if nothing was found and false if i is out of bounds.
func (l *LinkedList[T]) GetNode(i int) (node *Node[T], success bool) {
	if i == l.Size-1 {
		return l.Tail, true
	}
	curNode := l.Head
	for j := 0; j < i; j++ {
		if curNode == nil {
			return nil, false
		}
		curNode = curNode.Next
	}
	if curNode == nil {
		return nil, false
	}
	return curNode, true
}

// Pop removes the head of the list and returns its value.
//
// Returns false if there was no head
func (l *LinkedList[T]) Pop() (val T, success bool) {
	var ret T
	if l.Head == nil {
		return ret, false
	}

	oldHead := l.Head
	ret = oldHead.Value.(T)
	l.Head = l.Head.Next
	oldHead = nil
	l.Size--

	if l.Head == nil {
		l.Tail = nil
	}

	return ret, true
}

// Push inserts the value at the end of the list as new tail.
//
// Returns the inserted Node
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
// Returns the newely created Node
// Returns false if i is out of bounds.
func (l *LinkedList[T]) Insert(val T, i int) (node Node[T], success bool) {
	return Node[T]{nil, nil}, false
}

// Remove the Node at i.
//
// Returns the value at i, nil if i is out of bounds and true if a Node at i could be found and deleted.
func (l *LinkedList[T]) Remove(i int) (val T, success bool) {
	var value T
	return value, false
}
