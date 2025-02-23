package hashMap

import (
	linkedList "dsa/datastructures/linkedList/linkedListHM"
	"dsa/util/sugar"
	"reflect"
	"testing"
)

// Hint for this function: linked lists must be constructed backwards, because the push function is always inserting at head.
// This is done to avoid needing a tail without losing performance.
// Ordering is completely irrelevant for a hashmap.
func ll[K any, V any](keys []K, vals []V) *linkedList.LinkedList[K, V] {
	// Hint: if valls is nil then len(vals) == 0 is true
	// This means, the vals == nil check MUST be before the check for an empty list
	if keys == nil {
		var key K
		var val V
		return &linkedList.LinkedList[K, V]{
			Head: &linkedList.Node[K, V]{key, val, nil, nil},
			Size: 1,
		}
	} else if len(keys) == 0 {
		return &linkedList.LinkedList[K, V]{Head: nil, Size: 0}
	}

	node := &linkedList.Node[K, V]{keys[len(keys)-1], vals[len(vals)-1], nil, nil}
	ll := &linkedList.LinkedList[K, V]{node, uint(len(keys))}

	for i := len(keys) - 2; i > -1; i-- {
		newNode := &linkedList.Node[K, V]{keys[i], vals[i], nil, node}
		node.Next = newNode
		node = newNode
	}

	return ll
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
	if len(countMap) != 0 {
		return false
	}
	return true
}

func TestHashMap_Clear(t *testing.T) {
	type testCase[K any, V any] struct {
		name   string
		hm     HashMap[K, V]
		wantHM HashMap[K, V]
	}
	tests := []testCase[int, int]{
		{
			"empty map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
		},
		{
			"filled map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{1}),
				ll[int, int]([]int{2, 3}, []int{2, 3}),
			}, 3},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
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
	type args struct {
		key int
	}
	type testCase[K any, V any] struct {
		name      string
		hm        HashMap[K, V]
		args      args
		wantHM    HashMap[K, V]
		wantFound bool
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			args{1},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			false,
		},
		{
			"not found",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{4},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			false,
		},
		{
			"found",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{1},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
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
			if gotFound := tt.hm.ContainsKey(tt.args.key); gotFound != tt.wantFound {
				t.Errorf("ContainsKey() = %v, wantFound %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("ContainsKey(), hm %v, wantHM %v", tt.hm, tt.wantHM)
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
		hm        HashMap[K, V]
		args      args
		wantHM    HashMap[K, V]
		wantFound bool
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			args{1},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			false,
		},
		{
			"not found",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{2},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			false,
		},
		{
			"found",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{6},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
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
	type args struct {
		key int
	}
	type testCase[K any, V any] struct {
		name      string
		hm        HashMap[K, V]
		args      args
		wantHM    HashMap[K, V]
		wantVal   V
		wantFound bool
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			args{1},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			0,
			false,
		},
		{
			"not found",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{4},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			0,
			false,
		},
		{
			"found",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{3},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			6,
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotVal, gotFound := tt.hm.Get(tt.args.key)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Get() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Get() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("Get(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
		})
	}
}

func TestHashMap_GetKey(t *testing.T) {
	type args struct {
		value int
	}
	type testCase[K any, V any] struct {
		name      string
		hm        HashMap[K, V]
		args      args
		wantHM    HashMap[K, V]
		wantKey   K
		wantFound bool
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			args{1},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			0,
			false,
		},
		{
			"not found",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{2},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			0,
			false,
		},
		{
			"found",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			args{6},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			6,
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotKey, gotFound := tt.hm.GetKey(tt.args.value)
			if !reflect.DeepEqual(gotKey, tt.wantKey) {
				t.Errorf("GetKey() gotKey = %v, want %v", gotKey, tt.wantKey)
			}
			if gotFound != tt.wantFound {
				t.Errorf("GetKey() gotFound = %v, want %v", gotFound, tt.wantFound)
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
		hm       HashMap[K, V]
		wantHM   HashMap[K, V]
		wantBool bool
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			true,
		},
		{
			"filled map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
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
		hm       HashMap[K, V]
		wantHM   HashMap[K, V]
		wantKeys []K
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			make([]int, 0),
		},
		{
			"filled map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
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
		hm         HashMap[K, V]
		wantHM     HashMap[K, V]
		wantValues []V
	}
	tests := []testCase[int, int]{
		{
			"emtpy map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			make([]int, 0),
		},
		{
			"filled map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{4}),
				ll[int, int]([]int{2, 3}, []int{5, 6}),
			}, 3},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
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
		wantHM HashMap[int, int]
	}{
		{
			"size 0",
			args{0},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
		},
		{
			"size 1",
			args{1},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{2}),
			}, 1},
		},
		{
			"size 10",
			args{10},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}),
				ll[int, int]([]int{6, 7, 8, 9, 10}, []int{6, 7, 8, 9, 10}),
			}, 10},
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
	type args struct {
		key int
		val int
	}
	type testCase[K any, V any] struct {
		name          string
		hm            HashMap[K, V]
		args          args
		wantHM        HashMap[K, V]
		wantPairsSize int
	}
	tests := []testCase[int, int]{
		{
			"into empty map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{}, 0},
			args{1, 2},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{2}),
			}, 1},
			1,
		},
		{
			"into map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1}, []int{2}),
			}, 1},
			args{3, 4},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1, 3}, []int{2, 4}),
			}, 2},
			2,
		},
		{
			"resize map",
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1, 2, 3}, []int{1, 2, 3}),
			}, 3},
			args{4, 4},
			HashMap[int, int]{[]*linkedList.LinkedList[int, int]{
				ll[int, int]([]int{1, 2, 3, 4}, []int{1, 2, 3, 4}),
			}, 4},
			6, // we double the array size if it is full, so we need to get a len of 6 here.
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			tt.hm.Insert(tt.args.key, tt.args.val)
			if !reflect.DeepEqual(tt.hm, tt.wantHM) {
				t.Errorf("Insert(), hm %v, wantHM %v", tt.hm, tt.wantHM)
			}
			if len(tt.hm.Pairs) != tt.wantPairsSize {
				t.Errorf("Insert(), gotPairsSize %v, wantPairsSize %v", len(tt.hm.Pairs), tt.wantPairsSize)
			}
		})
	}
}

func TestHashMap_Remove(t *testing.T) {
	type args struct {
		key int
	}
	type testCase[K any, V any] struct {
		name      string
		hm        HashMap[K, V]
		args      args
		wantHM    HashMap[K, V]
		wantVal   V
		wantFound bool
	}
	tests := []testCase[ /*TODO: finish this test*/ ]{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		println()

		// TODO: Make sure the resizing is also tested!!! (if len(Pairs) >> 2 == HashMap.Size, we want to resize the map to have Pairs be double the size of HashMap.Size
		//		 This is just my approach to do it. no general rule. Resizing happens effectively at a load factor of 0.25 to be 0.5 afterwards.)
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotVal, gotFound := tt.hm.Remove(tt.args.key)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Remove() gotVal = %v, want %v", gotVal, tt.wantVal)
			}
			if gotFound != tt.wantFound {
				t.Errorf("Remove() gotFound = %v, want %v", gotFound, tt.wantFound)
			}
		})
	}
}
