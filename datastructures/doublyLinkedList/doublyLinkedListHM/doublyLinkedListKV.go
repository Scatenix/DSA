package doublyLinkedListHM

/*
This is an adjusted doubly linked list wit key and value field for hash maps by altering it to fit better in a hash map.

For the future: It would probably be better if Key and Value are pointers and all the returns of Key or Value as well.
*/

import (
	"reflect"
)

type Node[K any, V any] struct {
	Key   K
	Value V
	Next  *Node[K, V]
	Prev  *Node[K, V]
}

type LinkedList[K any, V any] struct {
	Head *Node[K, V]
	Size uint
}

// NewLinkedList creates a new linkedList with n elements of type T.
func NewLinkedList[K any, V any](key K, value V) *LinkedList[K, V] {
	ll := LinkedList[K, V]{Head: nil, Size: 0}
	ll.Push(key, value)
	return &ll
}

// IsEmpty returns true if the linkedList does not contain any nodes.
func (l *LinkedList[K, V]) IsEmpty() bool {
	if l.Size <= 0 {
		return true
	}
	return false
}

// GetNode which contains key.
//
// Returns the found node if key exists in linked list and whether a node was found.
func (l *LinkedList[K, V]) GetNode(key K) (node *Node[K, V], found bool) {
	if l.IsEmpty() {
		return &Node[K, V]{}, false
	}

	curNode := l.Head
	for curNode != nil {
		if reflect.DeepEqual(curNode.Key, key) {
			return curNode, true
		}
		curNode = curNode.Next
	}

	return &Node[K, V]{}, false
}

// GetNode which contains value.
//
// Returns the found node if value exists in linked list and whether a node was found.
func (l *LinkedList[K, V]) GetNodeByValue(value V) (node *Node[K, V], found bool) {
	if l.IsEmpty() {
		return &Node[K, V]{}, false
	}

	curNode := l.Head
	for curNode != nil {
		// Note: DeepEqual may not be able to compare some data types like time.Time (as of 2022).
		// It is recommended to use the "github.com/google/go-cmp/cmp" dependency.
		// In my case I won't do it to avoid getting external dependencies into this project.
		if reflect.DeepEqual(curNode.Value, value) {
			return curNode, true
		}
		curNode = curNode.Next
	}

	return &Node[K, V]{}, false
}

// Push inserts the value to the list.
//
// Returns the pushed Node.
//
// Note: This is actually a PushFront function.
// To avoid needing a tail and without losing performance, the key/value pair is always pushed as a new head,
// because ordering does not matter here.
func (l *LinkedList[K, V]) Push(key K, value V) *Node[K, V] {
	newNode := &Node[K, V]{key, value, l.Head, nil}

	l.Size++
	if l.Head != nil {
		l.Head.Prev = newNode
	}
	l.Head = newNode

	return newNode
}

// Remove the Node with key.
//
// Returns the removed node and whether a node with key could be found and removed.
//
// Returns an error if list is empty or i >= l.Size.
func (l *LinkedList[K, V]) Remove(key K) (v V, found bool) {
	if l.IsEmpty() {
		return v, false
	}

	delNode, exists := l.GetNode(key)
	if !exists {
		return v, false
	}

	l.Size--
	prevNode := delNode.Prev
	nextNode := delNode.Next
	if prevNode != nil {
		prevNode.Next = nextNode
		if nextNode != nil {
			nextNode.Prev = prevNode
		}
	} else {
		l.Head = nextNode
		if l.Head != nil {
			l.Head.Prev = nil
		}
	}
	v = delNode.Value
	delNode = nil

	return v, true
}
