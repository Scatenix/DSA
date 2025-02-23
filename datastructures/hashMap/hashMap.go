package hashMap

import (
	linkedList "dsa/datastructures/linkedList/linkedListHM"
)

/*
Dynamic resizing:
	Resizing is typically done by doubling the size of the underlying array and recalculating every key.
	TODO: Think about adding a retrival of all nodes per linked list as function within the linked list itself or do it in this hashMap with node traversal.
*/

// hash.Murmur3("test", 2)

type HashMap[K, V any] struct {
	Pairs []*linkedList.LinkedList[K, V]
	Size  uint
}

// NewHashMap creates a new HashMap.
//
// initialCapacity - The initial capacity of the HashMap on its creation.
func NewHashMap[K, V any](initialCapacity uint) *HashMap[K, V] {
	return &HashMap[K, V]{}
}

// Insert a key value pair.
func (hm *HashMap[K, V]) Insert(key K, val V) {
	// TODO: if len(Pairs) == HashMap.Size, we want to resize the map with double the current size of Pairs.
}

// Get value by key and whether any value could be found.
func (hm *HashMap[K, V]) Get(key K) (val V, found bool) {
	return val, found
}

// GetKey by value and whether any key could be found.
func (hm *HashMap[K, V]) GetKey(value V) (key K, found bool) {
	return key, found
}

// Remove key value pair by key.
//
// returns the value and wheter any pair was found to be removed.
func (hm *HashMap[K, V]) Remove(key K) (val V, found bool) {
	// TODO: if len(Pairs) >> 2 == HashMap.Size, we want to resize the map to have Pairs be double the size of HashMap.Size
	// 		 This is just my approach to do it. no general rule. Resizing happens effectively at a load factor of 0.25 to be 0.5 afterwards.
	//	Put this as a normal comment after implementation.
	return val, found
}

// ContainsKey - Check if key exists.
func (hm *HashMap[K, V]) ContainsKey(key K) bool {
	return false
}

// ContainsVal - Check if value exists.
func (hm *HashMap[K, V]) ContainsVal(val V) bool {
	return false
}

// IsEmpty - Check if map is emtpy.
func (hm *HashMap[K, V]) IsEmpty() bool {
	return true
}

// Keys returns an array of all keys.
func (hm *HashMap[K, V]) Keys() []K {
	return make([]K, 0)
}

// Values returns an array of all values.
func (hm *HashMap[K, V]) Values() []V {
	return make([]V, 0)
}

// Clear HashMap, resetting it to a newly initialized state.
func (hm *HashMap[K, V]) Clear() {
}
