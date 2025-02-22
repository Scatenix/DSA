package linkedList

import (
	"dsa/util/sugar"
	"reflect"
	"testing"
)

// Hint for this function: linked lists must be constructed backwards, because the push function is always inserting at head.
// This is done to avoid needing a tail without losing performance.
// Ordering is completely irrelevant for a hashmap.
func ll[K any, V any](keys []K, vals []V) *LinkedList[K, V] {
	// Hint: if valls is nil then len(vals) == 0 is true
	// This means, the vals == nil check MUST be before the check for an empty list
	if keys == nil {
		var key K
		var val V
		return &LinkedList[K, V]{
			Head: &Node[K, V]{key, val, nil, nil},
			Size: 1,
		}
	} else if len(keys) == 0 {
		return &LinkedList[K, V]{Head: nil, Size: 0}
	}

	node := &Node[K, V]{keys[len(keys)-1], vals[len(vals)-1], nil, nil}
	ll := &LinkedList[K, V]{node, uint(len(keys))}

	for i := len(keys) - 2; i > -1; i-- {
		newNode := &Node[K, V]{keys[i], vals[i], nil, node}
		node.Next = newNode
		node = newNode
	}

	return ll
}

func TestLinkedList_GetNode(t *testing.T) {
	type args[K any] struct {
		key K
	}
	type testCase[K any, V any] struct {
		name               string
		existingLinkedList *LinkedList[K, V]
		args               args[K]
		wantLL             *LinkedList[K, V]
		wantNode           *Node[K, V]
		wantFound          bool
	}
	tests := []testCase[int, int]{
		{
			"emtpy linkedList",
			ll[int, int]([]int{}, []int{}),
			args[int]{3},
			ll[int, int]([]int{}, []int{}),
			&Node[int, int]{},
			false,
		},
		{
			"populated linkedList, key does not exist",
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			args[int]{6},
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			&Node[int, int]{},
			false,
		},
		// This test case is cheated, because an linkedList[int] cannot have a Node with a value of nil,
		// but for this test case, it still gets the point.
		{
			"head value is nil",
			ll[int, int]([]int{}, []int{}),
			args[int]{0},
			ll[int, int]([]int{}, []int{}),
			&Node[int, int]{},
			false,
		},
		{
			"populated linkedList key at beginning",
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			args[int]{1},
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}).Head.Next.Next.Next.Next, // ll for my hashmap is backwards
			true,
		},
		{
			"populated linkedList key at end",
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			args[int]{5},
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}).Head, // ll for my hashmap is backwards
			true,
		},
		{
			"populated linkedList beginning < key < end",
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			args[int]{3},
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}).Head.Next.Next,
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotNode, gotFound := tt.existingLinkedList.GetNode(tt.args.key)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("actuall LL = %v, wantLL %v", tt.existingLinkedList, tt.wantLL)
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("GetNode() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotFound, tt.wantFound) {
				t.Errorf("GetNode() gotFound = %v, wantFound %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestLinkedList_GetNodeByValue(t *testing.T) {
	type args[V any] struct {
		value V
	}
	type testCase[K any, V any] struct {
		name               string
		existingLinkedList *LinkedList[K, V]
		args               args[V]
		wantLL             *LinkedList[K, V]
		wantNode           *Node[K, V]
		wantFound          bool
	}
	tests := []testCase[int, int]{
		{
			"emtpy linkedList",
			ll[int, int]([]int{}, []int{}),
			args[int]{3},
			ll[int, int]([]int{}, []int{}),
			&Node[int, int]{},
			false,
		},
		{
			"populated linkedList, value does not exist",
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			args[int]{5},
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			&Node[int, int]{},
			false,
		},
		// This test case is cheated, because an linkedList[int] cannot have a Node with a value of nil,
		// but for this test case, it still gets the point.
		{
			"head value is nil",
			ll[int, int]([]int{}, []int{}),
			args[int]{0},
			ll[int, int]([]int{}, []int{}),
			&Node[int, int]{},
			false,
		},
		{
			"populated linkedList value at beginning",
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			args[int]{6},
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}).Head.Next.Next.Next.Next, // ll for my hashmap is backwards
			true,
		},
		{
			"populated linkedList value at end",
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			args[int]{10},
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}).Head, // ll for my hashmap is backwards
			true,
		},
		{
			"populated linkedList beginning < key < end",
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			args[int]{8},
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}),
			ll[int, int]([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9, 10}).Head.Next.Next,
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotNode, gotFound := tt.existingLinkedList.GetNodeByValue(tt.args.value)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("actuall LL = %v, wantLL %v", tt.existingLinkedList, tt.wantLL)
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("GetNode() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotFound, tt.wantFound) {
				t.Errorf("GetNode() gotFound = %v, wantFound %v", gotFound, tt.wantFound)
			}
		})
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	type testCase[K, V any] struct {
		name               string
		existingLinkedList *LinkedList[K, V]
		wantLL             *LinkedList[K, V]
		wantBool           bool
	}
	tests := []testCase[int, int]{
		{
			"empty linkedList",
			ll[int, int]([]int{}, []int{}),
			ll[int, int]([]int{}, []int{}),
			true,
		},
		{
			"head value is nil",
			ll[int, int](nil, nil),
			ll[int, int](nil, nil),
			false,
		},
		{
			"populated linkedList",
			ll[int, int]([]int{1, 2, 3}, []int{1, 2, 3}),
			ll[int]([]int{1, 2, 3}, []int{1, 2, 3}),
			false,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("actuall LL = %v, wantLL %v", tt.existingLinkedList, tt.wantLL)
			}
			if gotBool := tt.existingLinkedList.IsEmpty(); gotBool != tt.wantBool {
				t.Errorf("IsEmpty() = %v, wantBool %v", gotBool, tt.wantBool)
			}
		})
	}
}

func TestLinkedList_Push(t *testing.T) {
	type args[K, V any] struct {
		key K
		val V
	}
	type testCase[K, V any] struct {
		name               string
		existingLinkedList *LinkedList[K, V]
		args               args[K, V]
		wantLL             *LinkedList[K, V]
		wantNode           *Node[K, V]
	}

	tests := []testCase[int, int]{
		{
			"int, int: push int to empty ll",
			ll[int, int]([]int{}, []int{}),
			args[int, int]{1, 2},
			ll[int, int]([]int{1}, []int{2}),
			ll[int, int]([]int{1}, []int{2}).Head,
		},
		{
			"int, int: push int to existing ll",
			ll[int]([]int{1}, []int{1}),
			args[int, int]{2, 3},
			ll[int]([]int{1, 2}, []int{1, 3}),
			ll[int]([]int{1, 2}, []int{1, 3}).Head,
		},
	}

	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotNode := tt.existingLinkedList.Push(tt.args.key, tt.args.val)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("NewLinkedList() = %v, wantVal %v", tt.existingLinkedList, tt.wantLL)
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("Get() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
		})
	}
}

func TestNewLinkedList(t *testing.T) {
	type args[K, V any] struct {
		key K
		val V
	}
	type testCase[K, V any] struct {
		name   string
		args   args[K, V]
		wantLL *LinkedList[K, V]
	}
	tests := []testCase[int, int]{
		{
			"create new list",
			args[int, int]{0, 1},
			ll[int]([]int{0}, []int{1}),
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			//t.Skip()
			if got := NewLinkedList(tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.wantLL) {
				t.Errorf("NewLinkedList() = %v, wantLL %v", got, tt.wantLL)
			}
		})
	}
}

func TestLinkedList_Remove(t *testing.T) {
	type args[K any] struct {
		key K
	}
	type testCase[K, V any] struct {
		name               string
		existingLinkedList *LinkedList[K, V]
		args               args[K]
		wantLL             *LinkedList[K, V]
		wantVal            V
		wantFound          bool
	}
	tests := []testCase[int, int]{
		{
			"remove from empty linkedList",
			ll[int, int]([]int{}, []int{}),
			args[int]{0},
			ll[int, int]([]int{}, []int{}),
			0,
			false,
		},
		{
			"remove not existing key",
			ll[int, int]([]int{1}, []int{1}),
			args[int]{3},
			ll[int, int]([]int{1}, []int{1}),
			0,
			false,
		},
		{
			"remove only item of list",
			ll[int, int]([]int{1}, []int{2}),
			args[int]{1},
			ll[int, int]([]int{}, []int{}),
			2,
			true,
		},
		{
			"remove head of list",
			ll[int, int]([]int{1, 2, 3}, []int{4, 5, 6}),
			args[int]{1},
			ll[int, int]([]int{2, 3}, []int{5, 6}),
			4,
			true,
		},
		{
			"remove tail of list",
			ll[int, int]([]int{1, 2, 3}, []int{4, 5, 6}),
			args[int]{3},
			ll[int, int]([]int{1, 2}, []int{4, 5}),
			6,
			true,
		},
		{
			"remove middle of list",
			ll[int, int]([]int{1, 2, 3}, []int{4, 5, 6}),
			args[int]{2},
			ll[int, int]([]int{1, 3}, []int{4, 6}),
			5,
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotVal, gotFound := tt.existingLinkedList.Remove(tt.args.key)
			if !reflect.DeepEqual(gotFound, tt.wantFound) {
				t.Errorf("Remove() gotFound = %v, wantFound %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Remove() gotVal = %v, wantVal %v", gotVal, tt.wantVal)
			}
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("actual LL = %v, wantLL %v", tt.existingLinkedList, tt.wantLL)
			}
		})
	}
}
