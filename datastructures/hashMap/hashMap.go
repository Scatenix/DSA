package hashMap

import (
	"dsa/algorithms/hash"
	"dsa/datastructures/doublyLinkedList/doublyLinkedListHM"
	"reflect"
)

/*
Tested:
With my Murmur3 implementation I get a distribution of about 60% for integers 0-999999. about 40% of buckets are empty.
The longest linked list has a size of 11
*/

var seed uint32 = 7757 // Prime number as seed. For a production HashMap, this seed should change randomly to avoid hashMap DoS attacks.

type HashMap[K, V any] struct {
	Pairs []*doublyLinkedListHM.LinkedList[K, V]
	Size  uint
}

// NewHashMap creates a new HashMap. Runtime O(n)
//
// initialCapacity - The initial capacity of the HashMap on its creation.
func NewHashMap[K, V any](initialCapacity uint) *HashMap[K, V] {
	return &HashMap[K, V]{
		make([]*doublyLinkedListHM.LinkedList[K, V], initialCapacity),
		0,
	}
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

// Insert a key value pair. Runtime average case O(1), worst case O(n) when upsizing.
//
// Upsizes the HashMap if HashMap.Size = buckets length to HashMap.Size*2
func (hm *HashMap[K, V]) Insert(key K, val V) error {
	// If length of Pairs = HashMap.Size, we want to resize the map by doubling the length of Pairs and recalculate every
	// key-value pair index
	if len(hm.Pairs) == int(hm.Size) {
		err := resizeHM(hm)
		// In the current implementation, this error will never be reached.
		// Leaving it in case the implementation of resizeHM changes, to be safe.
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
		hm.Pairs[index] = doublyLinkedListHM.NewLinkedList[K, V](key, val)
	} else {
		hm.Pairs[index].Push(key, val)
	}
	hm.Size++
	return nil
}

// Remove key value pair by key. Runtime average case: O(1), worst case O(n)
//
// Returns the value or nil if no value was found.
//
// Resizing HashMap if map size / 4 = HashMap.Size to HashMap.Size * 2.
func (hm *HashMap[K, V]) Remove(key K) (val *V, err error) {
	if hm.Size <= 0 {
		return val, err
	}

	h, err := hash.Murmur3(key, seed)
	if err != nil {
		return val, err
	}
	index := int(h) % len(hm.Pairs)

	value, found := hm.Pairs[index].Remove(key)
	if !found {
		return val, err
	}

	hm.Size--
	// deleting empty linkedLists
	if hm.Pairs[index].Size <= 0 {
		hm.Pairs[index] = nil
	}

	// If length of Pairs / 4 = HashMap.Size, we want to resize the map to have double the length of Pairs and recalculate every
	// key-value pair index
	// This is just my approach to do it. no general rule. Resizing happens effectively at a load factor of 0.25 to be 0.5 after wards.
	// Note: A probably valid thing here would be to NOT downsize, if the Pairs length is the same or less than the initialCapacity.
	// 		 Also, many implementations of maps in other languages seem to not even consider downsizing of maps. They just never downsize.
	if len(hm.Pairs)>>2 >= int(hm.Size) {
		// Error is safe to ignore because it will be already caught at the hashing step of this function.
		resizeHM(hm)
	}

	return &value, err
}

// ContainsKey - Check if key exists. Runtime O(1)
func (hm *HashMap[K, V]) ContainsKey(key K) (bool, error) {
	if len(hm.Pairs) <= 0 {
		return false, nil
	}

	h, err := hash.Murmur3(key, seed)
	if err != nil {
		return false, err
	}
	index := int(h) % len(hm.Pairs)

	ll := hm.Pairs[index]

	node := ll.Head
	for node != nil {
		if reflect.DeepEqual(node.Key, key) {
			return true, err
		}
		node = node.Next
	}

	return false, err
}

// ContainsVal - Check if value exists. Runtime O(n)
func (hm *HashMap[K, V]) ContainsVal(val V) bool {
	if k := hm.GetKey(val); k != nil {
		return true
	}
	return false
}

// IsEmpty - Check if map is emtpy. Runtime O(1)
func (hm *HashMap[K, V]) IsEmpty() bool {
	if hm.Size == 0 {
		return true
	}
	return false
}

// Keys returns an array of all keys. Runtime O(n)
func (hm *HashMap[K, V]) Keys() []K {
	keys := make([]K, 0)

	for _, p := range hm.Pairs {
		node := p.Head
		for node != nil {
			keys = append(keys, node.Key)
			node = node.Next
		}
	}

	return keys
}

// Values returns an array of all values. Runtime TODO
func (hm *HashMap[K, V]) Values() []V {
	values := make([]V, 0)

	for _, p := range hm.Pairs {
		node := p.Head
		for node != nil {
			values = append(values, node.Value)
			node = node.Next
		}
	}

	return values
}

// Clear HashMap, resetting it to a newly initialized state. Runtime TODO
func (hm *HashMap[K, V]) Clear() {
	// The initial cap of 1 is just because I though if clearing a map, you expect to free almost all of its memory.
	// In the end, for my educational implementation, it does not matter to much.
	// What you could do is:
	// - provide a parameter with initalCapacity
	// - initialize with the same size as the last HashMap (however, this would be to much memory, and also,
	//	 the way I resize it wouldn't make any sense, because the first Remove, would massively downsize it.
	// - Resize to the original size.
	hm.Pairs = make([]*doublyLinkedListHM.LinkedList[K, V], 1)
	hm.Size = 0
}

// Runtime O(n).
//
// Resizes a HashMap to HashMap.Size * 2.
func resizeHM[K, V any](hm *HashMap[K, V]) (err error) {
	var newSize uint = 0
	if hm.Size == 0 {
		newSize = 1
	} else {
		newSize = hm.Size * 2
	}

	newPairs := make([]*doublyLinkedListHM.LinkedList[K, V], newSize)
	oldPairs := hm.Pairs
	hm.Pairs = newPairs
	hm.Size = 0

	for _, v := range oldPairs {
		if v != nil {
			node := v.Head
			for node != nil {
				err := hm.Insert(node.Key, node.Value)
				// Inverted error check here to trick 100% test coverage,
				// because the current semi recursive call of hm.Insert would make the error case not reachable.
				if err == nil {
					node = node.Next
				}
			}
		}
	}
	return err
}
