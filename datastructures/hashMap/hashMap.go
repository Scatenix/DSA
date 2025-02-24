package hashMap

import (
	"dsa/algorithms/hash"
	linkedList "dsa/datastructures/linkedList/linkedListHM"
	"reflect"
)

/*
Dynamic resizing:
	Resizing is typically done by doubling the size of the underlying array and recalculating every key.
	TODO: Think about adding a retrival of all nodes per linked list as function within the linked list itself or do it in this hashMap with node traversal.
*/

var seed uint32 = 7757 // Prime number as seed. For a production HashMap, this seed should change randomly to avoid hashMap ddos attacks.

type HashMap[K, V any] struct {
	Pairs []*linkedList.LinkedList[K, V]
	Size  uint
}

// NewHashMap creates a new HashMap. Runtime O(n)
//
// initialCapacity - The initial capacity of the HashMap on its creation.
func NewHashMap[K, V any](initialCapacity uint) *HashMap[K, V] {
	return &HashMap[K, V]{
		make([]*linkedList.LinkedList[K, V], initialCapacity),
		0,
	}
}

// Insert a key value pair. Runtime average case O(1), worst case O(n) when upsizing.
//
// Upsizes the HashMap if HashMap.Size = buckets length to HashMap.Size*2
func (hm *HashMap[K, V]) Insert(key K, val V) error {
	// If length of Pairs = HashMap.Size, we want to resize the map by doubling the length of Pairs and recalculate every
	// key-value pair index
	if len(hm.Pairs) == int(hm.Size) {
		err := upsizeHM(hm)
		if err != nil {
			return err
		}
	}

	h, err := hash.Murmur3(key, seed)
	if err != nil {
		return err
	}

	index := int(h) % len(hm.Pairs)

	if hm.Pairs[index] == nil {
		hm.Pairs[index] = linkedList.NewLinkedList[K, V](key, val)
	} else {
		hm.Pairs[index].Push(key, val)
	}
	hm.Size++
	return nil
}

// Get value by key. Runtime O(1)
//
// Returns nil if no value was found
func (hm *HashMap[K, V]) Get(key K) (val *V, err error) {
	if len(hm.Pairs) == 0 {
		return val, nil
	}

	h, err := hash.Murmur3(key, seed)
	if err != nil {
		return val, err
	}
	index := int(h) % len(hm.Pairs)

	ll := hm.Pairs[index]

	node := ll.Head
	for node != nil {
		if reflect.DeepEqual(node.Key, key) {
			return &node.Value, err
		}
		node = node.Next
	}

	return val, err
}

// GetKey by value and whether any key could be found. Runtime O(n)
//
// # Returns nil if no key was found
func (hm *HashMap[K, V]) GetKey(value V) (key *K) {
	for _, v := range hm.Pairs {
		if v != nil {
			node := v.Head
			for node != nil {
				if reflect.DeepEqual(node.Value, value) {
					return &node.Key
				}
				node = node.Next
			}
		}
	}

	return key
}

// Remove key value pair by key. Runtime TODO
//
// returns the value and wheter any pair was found to be removed.
func (hm *HashMap[K, V]) Remove(key K) (val *V) {
	// TODO: if len(Pairs) >> 2 == HashMap.Size, we want to resize the map to have Pairs be double the size of HashMap.Size
	// 		 This is just my approach to do it. no general rule. Resizing happens effectively at a load factor of 0.25 to be 0.5 afterwards.
	//	Put this as a normal comment after implementation.
	return val
}

// ContainsKey - Check if key exists. Runtime TODO
func (hm *HashMap[K, V]) ContainsKey(key K) bool {
	return false
}

// ContainsVal - Check if value exists. Runtime TODO
func (hm *HashMap[K, V]) ContainsVal(val V) bool {
	return false
}

// IsEmpty - Check if map is emtpy. Runtime TODO
func (hm *HashMap[K, V]) IsEmpty() bool {
	return true
}

// Keys returns an array of all keys. Runtime TODO
func (hm *HashMap[K, V]) Keys() []K {
	return make([]K, 0)
}

// Values returns an array of all values. Runtime TODO
func (hm *HashMap[K, V]) Values() []V {
	return make([]V, 0)
}

// Clear HashMap, resetting it to a newly initialized state. Runtime TODO
func (hm *HashMap[K, V]) Clear() {
}

// Runtime O(n)
func upsizeHM[K, V any](hm *HashMap[K, V]) error {
	var newSize uint = 0
	if hm.Size == 0 {
		newSize = 1
	} else {
		newSize = hm.Size * 2
	}

	newPairs := make([]*linkedList.LinkedList[K, V], newSize)
	oldPairs := hm.Pairs
	hm.Pairs = newPairs
	hm.Size = 0

	for _, v := range oldPairs {
		if v != nil {
			node := v.Head
			for node != nil {
				err := hm.Insert(node.Key, node.Value)
				if err != nil {
					return err
				}
				node = node.Next
			}
		}
	}
	return nil
}
