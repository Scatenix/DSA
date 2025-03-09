package hashMap

import (
	"dsa/datastructures/doublyLinkedList/doublyLinkedListHM"
	"dsa/util/sugar"
	"reflect"
	"testing"
)

/*
Note for all of these tests:
The linked list structures are not important at all for the HashMap.
We only want to know if all the key-value pairs are present, but do not care and because of the hashing cannot even know where exactly they are.
*/

// Hint for this function: linked lists must be constructed backwards, because the push function is always inserting at head.
// This is done to avoid needing a tail without losing performance.
// Ordering is completely irrelevant for a hashmap.
func ll[K any, V any](keys []K, vals []V) *doublyLinkedListHM.LinkedList[K, V] {
	// Hint: if valls is nil then len(vals) == 0 is true
	// This means, the vals == nil check MUST be before the check for an empty list
	if keys == nil {
		var key K
		var val V
		return &doublyLinkedListHM.LinkedList[K, V]{
			Head: &doublyLinkedListHM.Node[K, V]{key, val, nil, nil},
			Size: 1,
		}
	} else if len(keys) == 0 {
		return &doublyLinkedListHM.LinkedList[K, V]{Head: nil, Size: 0}
	}

	node := &doublyLinkedListHM.Node[K, V]{keys[len(keys)-1], vals[len(vals)-1], nil, nil}
	ll := &doublyLinkedListHM.LinkedList[K, V]{node, uint(len(keys))}

	for i := len(keys) - 2; i > -1; i-- {
		newNode := &doublyLinkedListHM.Node[K, V]{keys[i], vals[i], nil, node}
		node.Next = newNode
		node = newNode
	}

	return ll
}

func intP(i int) *int {
	return &i
}

func slicesEqualUnordered(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	countMap := make(map[int]int)

	for _, v := range a {
		countMap[v]++
	}
	for _, v := range b {
		countMap[v]--
		if countMap[v] < 0 {
			return false
		}
	}
	for _, v := range countMap {
		if v != 0 {
			return false
		}
	}
	return true
}

// mapsEqual only needed for tests that resize the map because the order and structure of the linked lists will be different.
func mapsEqual(a, b *HashMap[int, int]) bool {
	if a.Size != b.Size {
		return false
	}
	aKeys := make([]int, 0)
	aVales := make([]int, 0)
	for _, v := range a.Pairs {
		if v != nil && !v.IsEmpty() {
			node := v.Head
			if node != nil {
				aKeys = append(aKeys, node.Key)
				aVales = append(aVales, node.Value)
				node = node.Next
			}
		}
	}

	bKeys := make([]int, 0)
	bVales := make([]int, 0)
	for _, v := range a.Pairs {
		if v != nil && !v.IsEmpty() {
			node := v.Head
			if node != nil {
				bKeys = append(bKeys, node.Key)
				bVales = append(bVales, node.Value)
				node = node.Next
			}
		}
	}

	if !slicesEqualUnordered(aKeys, bKeys) || !slicesEqualUnordered(aVales, bVales) {
		return false
	}
	return true
}

func TestHashMap_Clear(t *testing.T) {
	type testCase[K any, V any] struct {
		name   string
		hm     *HashMap[K, V]
		wantHM *HashMap[K, V]
	}
	tests := []testCase[int, int]{
		{
			"empty map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{nil}, 0}, // Putting nil in ll slice because I clear maps to have a length 1
		},
		{
			"filled map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{1}),
				ll[int, int]([]int{2, 3}, []int{2, 3}),
			}, 3},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{nil}, 0}, // Putting nil in ll slice because I clear maps to have a length 1
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			tt.hm.Clear()
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("Clear(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
		})
	}
}

func TestHashMap_ContainsKey(t *testing.T) {
	type args[K any] struct {
		key K
	}
	type testCase[K any, V any] struct {
		name      string
		hm        *HashMap[K, V]
		args      args[K]
		wantHM    *HashMap[K, V]
		wantFound bool
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			args[int]{1},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			false,
		},
		{
			"not found",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args[int]{4},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			false,
		},
		{
			"found",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args[int]{1},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			true,
		},
	}
	testError := []testCase[chan int, chan int]{
		{
			// My hashing function does not support values of type channel
			"test error",
			&HashMap[chan int, chan int]{[]*doublyLinkedListHM.LinkedList[chan int, chan int]{
				ll[chan int, chan int]([]chan int{}, []chan int{}),
			}, 1},
			args[chan int]{nil},
			nil, // We do not care for it in this test.
			false,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			if gotFound, _ := tt.hm.ContainsKey(tt.args.key); gotFound != tt.wantFound {
				t.Errorf("ContainsKey() = %v, wantFound %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("ContainsKey(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
		})
	}
	for _, tt := range testError {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			_, err := tt.hm.ContainsKey(tt.args.key)
			if err == nil {
				t.Errorf("test %v should have thrown error. Got %v", tt.name, err)
			}
		})
	}
}

func TestHashMap_ContainsVal(t *testing.T) {
	type args struct {
		val int
	}
	type testCase[K any, V any] struct {
		name      string
		hm        *HashMap[K, V]
		args      args
		wantHM    *HashMap[K, V]
		wantFound bool
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			args{1},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			false,
		},
		{
			"not found",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{2},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			false,
		},
		{
			"found",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{6},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			if got := tt.hm.ContainsVal(tt.args.val); got != tt.wantFound {
				t.Errorf("ContainsVal() = %v, wantFound %v", got, tt.wantFound)
			}
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("ContainsVal(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
		})
	}
}

func TestHashMap_Get(t *testing.T) {
	type args[T any] struct {
		key T
	}
	type testCase[K any, V any] struct {
		name    string
		hm      *HashMap[K, V]
		args    args[K]
		wantHM  *HashMap[K, V]
		wantVal *V
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			args[int]{1},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			nil,
		},
		{
			"not found",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args[int]{4},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			nil,
		},
		{
			"found",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1, 3}, []int{4, 6}), // hash algorithm will place k,v at index 0.
				ll[int, int]([]int{2}, []int{5}),
			}, 3},
			args[int]{3},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1, 3}, []int{4, 6}),
				ll[int, int]([]int{2}, []int{5}),
			}, 3},
			intP(6),
		},
	}
	testError := []testCase[chan int, chan int]{
		{
			// My hashing function does not support values of type channel
			"test error",
			&HashMap[chan int, chan int]{[]*doublyLinkedListHM.LinkedList[chan int, chan int]{
				ll[chan int, chan int]([]chan int{}, []chan int{}),
			}, 1},
			args[chan int]{nil},
			nil, // We do not care for it in this test.
			nil,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotVal, err := tt.hm.Get(tt.args.key)
			if err != nil {
				t.Errorf("Hash function threw error: %v", err)
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Get() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("Get(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
		})
	}
	for _, tt := range testError {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			_, err := tt.hm.Get(tt.args.key)
			if err == nil {
				t.Errorf("test %v should have thrown error. Got %v", tt.name, err)
			}
		})
	}
}

func TestHashMap_GetKey(t *testing.T) {
	type args struct {
		value int
	}
	type testCase[K any, V any] struct {
		name    string
		hm      *HashMap[K, V]
		args    args
		wantHM  *HashMap[K, V]
		wantKey *K
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			args{1},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			nil,
		},
		{
			"not found",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{2},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			nil,
		},
		{
			"found",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1, 3}, []int{4, 6}), // hash algorithm will place k,v at index 0.
				ll[int, int]([]int{2}, []int{5}),
			}, 3},
			args{6},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1, 3}, []int{4, 6}),
				ll[int, int]([]int{2}, []int{5}),
			}, 3},
			intP(3),
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotKey := tt.hm.GetKey(tt.args.value)
			if !reflect.DeepEqual(gotKey, tt.wantKey) {
				t.Errorf("GetKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("GetKey(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
		})
	}
}

func TestHashMap_IsEmpty(t *testing.T) {
	type testCase[K any, V any] struct {
		name     string
		hm       *HashMap[K, V]
		wantHM   *HashMap[K, V]
		wantBool bool
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			true,
		},
		{
			"filled map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			if got := tt.hm.IsEmpty(); got != tt.wantBool {
				t.Errorf("IsEmpty() = %v, wantBool %v", got, tt.wantBool)
			}
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("IsEmpty(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
		})
	}
}

func TestHashMap_Keys(t *testing.T) {
	type testCase[K any, V any] struct {
		name     string
		hm       *HashMap[K, V]
		wantHM   *HashMap[K, V]
		wantKeys []K
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			make([]int, 0),
		},
		{
			"filled map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			[]int{1, 2, 3},
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			if got := tt.hm.Keys(); !slicesEqualUnordered(got, tt.wantKeys) {
				t.Errorf("Keys() = %v, wantKeys %v", got, tt.wantKeys)
			}
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("Keys(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
		})
	}
}

func TestHashMap_Values(t *testing.T) {
	type testCase[K any, V any] struct {
		name       string
		hm         *HashMap[K, V]
		wantHM     *HashMap[K, V]
		wantValues []V
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			make([]int, 0),
		},
		{
			"filled map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			[]int{4, 5, 6},
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			if got := tt.hm.Values(); !slicesEqualUnordered(got, tt.wantValues) {
				t.Errorf("Values() = %v, wantValues %v", got, tt.wantValues)
			}
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("Values(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
		})
	}
}

func TestNewHashMap(t *testing.T) {
	type args struct {
		initialCapacity uint
	}
	tests := []struct {
		name   string
		args   args
		wantHM *HashMap[int, int]
	}{
		{
			"size 0",
			args{0},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
		},
		{
			"size 1",
			args{1},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{nil}, 0},
		},
		{
			"size 10",
			args{10},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				nil, nil, nil, nil, nil, nil, nil, nil, nil, nil,
			}, 0},
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			if got := NewHashMap[int, int](tt.args.initialCapacity); !reflect.DeepEqual(got, tt.wantHM) {
				t.Errorf("NewHashMap(), hm %v, wantHM %v", got, tt.wantHM)
			}
		})
	}
}

func TestHashMap_Insert(t *testing.T) {
	type args[K, V any] struct {
		key K
		val V
	}
	type testCase[K any, V any] struct {
		name          string
		hm            *HashMap[K, V]
		args          args[K, V]
		wantHM        *HashMap[K, V]
		wantPairsSize int
	}
	tests := []testCase[int, int]{
		{
			"into empty map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			args[int, int]{1, 2},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{2}),
			}, 1},
			1,
		},
		{
			"into map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{2}),
			}, 1},
			args[int, int]{3, 4},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1, 3}, []int{2, 4}),
			}, 2},
			2,
		},
		{
			"upsize map",
			// it is important for the test, that the Pairs length starts with 3
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{1}),
				ll[int, int]([]int{2, 3}, []int{2, 3}),
				nil,
			}, 3},
			args[int, int]{4, 4},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}),
			}, 4},
			6, // we double the array size if it is full, so we need to get a len of 6 here.
		},
	}
	testError := []testCase[chan int, chan int]{
		{
			// My hashing function does not support values of type channel
			"test error",
			&HashMap[chan int, chan int]{[]*doublyLinkedListHM.LinkedList[chan int, chan int]{
				ll[chan int, chan int]([]chan int{}, []chan int{}),
			}, 0},
			args[chan int, chan int]{nil, nil},
			nil, // We do not care for it in this test.
			0,   // We do not care for it in this test.
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			tt.hm.Insert(tt.args.key, tt.args.val)
			if !mapsEqual(tt.hm, tt.wantHM) {
				t.Errorf("Insert(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
			if len(tt.hm.Pairs) != tt.wantPairsSize {
				t.Errorf("Insert(), gotPairsSize %v, wantPairsSize %v", len(tt.hm.Pairs), tt.wantPairsSize)
			}
		})
	}
	for _, tt := range testError {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			err := tt.hm.Insert(tt.args.key, tt.args.val)
			if err == nil {
				t.Errorf("test %v should have thrown error. Got %v", tt.name, err)
			}
		})
	}
}

func TestHashMap_Remove(t *testing.T) {
	type args[K any] struct {
		key K
	}
	type testCase[K any, V any] struct {
		name    string
		hm      *HashMap[K, V]
		args    args[K]
		wantHM  *HashMap[K, V]
		wantVal *V
	}
	tests := []testCase[int, int]{
		{
			"empty map",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			args[int]{1},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			nil,
		},
		{
			"remove from map size 1",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{2}),
			}, 1},
			args[int]{1},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{}, 0},
			intP(2),
		},
		{
			"remove from map size 4 with ll.Size = 1",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1, 2}, []int{1, 2}),
				ll[int, int]([]int{4}, []int{4}),
				ll[int, int]([]int{3}, []int{3}), // need to place value here because hashing will place key 3 at index 2
			}, 4},
			args[int]{3},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1, 2}, []int{1, 2}),
				ll[int, int]([]int{4}, []int{4}),
			}, 3},
			intP(3),
		},
		{
			"remove from map size 4 with ll.Size > 1",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{3}, []int{3}),
				ll[int, int]([]int{4}, []int{4}),
				ll[int, int]([]int{1, 2}, []int{1, 2}), // need to place value here because hashing will place key 2 at index 2
			}, 4},
			args[int]{2},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{1}),
				ll[int, int]([]int{3}, []int{3}),
				ll[int, int]([]int{4}, []int{4}),
			}, 3},
			intP(2),
		},
		{
			"try remove not existing key",
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{2}),
			}, 1},
			args[int]{2},
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{2}),
			}, 1},
			nil,
		},
		{
			"downsize map",
			// Create an initial HashMap with a Paris array of size 20:
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{3, 4, 5}, []int{3, 4, 5}), // need to place value here because hashing will place key 5 at index 8
				ll[int, int]([]int{1, 2}, []int{1, 2}),
				nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, nil, // 18 empty buckets
			}, 5},
			args[int]{5},
			// Create a want to have HashMap with a Pairs array of size 8, because we are sizing down if len(Pairs) / 4 == HashMap.Size to len(Pairs) * 2
			&HashMap[int, int]{[]*doublyLinkedListHM.LinkedList[int, int]{
				ll[int, int]([]int{1, 2}, []int{1, 2}),
				ll[int, int]([]int{3, 4}, []int{3, 4}),
				nil, nil, nil, nil, nil, nil,
			}, 4},
			intP(5), // we double the array size if it is full, so we need to get a len of 6 here.
		},
	}
	testError := []testCase[chan int, chan int]{
		{
			// My hashing function does not support values of type channel
			"test error",
			&HashMap[chan int, chan int]{[]*doublyLinkedListHM.LinkedList[chan int, chan int]{
				ll[chan int, chan int]([]chan int{}, []chan int{}),
			}, 1},
			args[chan int]{nil},
			nil, // We do not care for it in this test.
			nil, // We do not care for it in this test.
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotVal, err := tt.hm.Remove(tt.args.key)
			if err != nil {
				t.Errorf("Hash function threw error: %v", err)
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Remove() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if !mapsEqual(tt.hm, tt.wantHM) {
				t.Errorf("Remove(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
		})
	}
	for _, tt := range testError {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			_, err := tt.hm.Remove(tt.args.key)
			if err == nil {
				t.Errorf("test %v should have thrown error. Got %v", tt.name, err)
			}
		})
	}
}
