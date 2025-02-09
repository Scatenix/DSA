package linkedList

import (
	"dsa/util/sugar"
	"errors"
	"reflect"
	"testing"
)

func ll[T any](vals []T) *LinkedList[T] {
	// Hint: if valls is nil then len(vals) == 0 is true
	// This means, the vals == nil check MUST be before the check for an empty list
	if vals == nil {
		return &LinkedList[T]{
			Head: &Node[T]{nil, nil},
			Tail: &Node[T]{nil, nil},
			Size: 1,
		}
	} else if len(vals) == 0 {
		return &LinkedList[T]{Head: nil, Tail: nil, Size: 0}
	}

	node := &Node[T]{vals[0], nil}
	ll := &LinkedList[T]{node, nil, uint(len(vals))}

	for i := 1; i < len(vals); i++ {
		newNode := &Node[T]{vals[i], nil}
		node.Next = newNode
		node = newNode
	}

	ll.Tail = node
	return ll
}

func node[T any](val T, next *Node[T]) *Node[T] {
	return &Node[T]{val, next}
}

func TestLinkedList_Get(t *testing.T) {
	type args struct {
		i uint
	}
	type testCase[T any] struct {
		name               string
		existingLinkedList *LinkedList[T]
		args               args
		wantLL             *LinkedList[T]
		wantVal            T
		wantBool           bool
	}
	ints := []int{1, 2, 3, 4, 5}
	tests := []testCase[*int]{
		{
			"emtpy linkedList",
			ll[*int]([]*int{}),
			args{3},
			ll[*int]([]*int{}),
			nil,
			false,
		},
		{
			"populated linkedList out of bounds",
			ll[*int]([]*int{&ints[0], &ints[1], &ints[2], &ints[3], &ints[4]}),
			args{10},
			ll[*int]([]*int{&ints[0], &ints[1], &ints[2], &ints[3], &ints[4]}),
			nil,
			false,
		},
		{
			"head value is nil",
			ll[*int]([]*int{nil}),
			args{0},
			ll[*int]([]*int{nil}),
			nil,
			true,
		},
		{
			"populated linkedList i at beginning",
			ll[*int]([]*int{&ints[0], &ints[1], &ints[2], &ints[3], &ints[4]}),
			args{0},
			ll[*int]([]*int{&ints[0], &ints[1], &ints[2], &ints[3], &ints[4]}),
			&ints[0],
			true,
		},
		{
			"populated linkedList i at end",
			ll[*int]([]*int{&ints[0], &ints[1], &ints[2], &ints[3], &ints[4]}),
			args{4},
			ll[*int]([]*int{&ints[0], &ints[1], &ints[2], &ints[3], &ints[4]}),
			&ints[4],
			true,
		},
		{
			"populated linkedList beginning < i < end ",
			ll[*int]([]*int{&ints[0], &ints[1], &ints[2], &ints[3], &ints[4]}),
			args{2},
			ll[*int]([]*int{&ints[0], &ints[1], &ints[2], &ints[3], &ints[4]}),
			&ints[2],
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotVal, gotBool := tt.existingLinkedList.Get(tt.args.i)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("Get() actual LL = %v, wantLL %v", tt.existingLinkedList, tt.wantLL)
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Get() gotVal = %v, wantVal %v", gotVal, tt.wantVal)
			}
			if gotBool != tt.wantBool {
				t.Errorf("Get() gotBool = %v, wantBool %v", gotBool, tt.wantBool)
			}
		})
	}
}

func TestLinkedList_GetNode(t *testing.T) {
	type args struct {
		i uint
	}
	type testCase[T any] struct {
		name               string
		existingLinkedList *LinkedList[T]
		args               args
		wantLL             *LinkedList[T]
		wantNode           *Node[T]
		wantBool           bool
	}
	tests := []testCase[int]{
		{
			"emtpy linkedList",
			ll[int]([]int{}),
			args{3},
			ll[int]([]int{}),
			nil,
			false,
		},
		{
			"populated linkedList out of bounds",
			ll[int]([]int{1, 2, 3, 4, 5}),
			args{10},
			ll[int]([]int{1, 2, 3, 4, 5}),
			nil,
			false,
		},
		{
			"populated linkedList out of bounds by one",
			ll[int]([]int{1, 2, 3, 4, 5}),
			args{5},
			ll[int]([]int{1, 2, 3, 4, 5}),
			nil,
			false,
		},
		// This test case is cheated, because an linkedList[int] cannot have a Node with a value of nil,
		// but for this test case, it still gets the point.
		{
			"head value is nil",
			ll[int](nil),
			args{0},
			ll[int](nil),
			&Node[int]{nil, nil},
			true,
		},
		{
			"populated linkedList i at beginning",
			ll[int]([]int{1, 2, 3, 4, 5}),
			args{0},
			ll[int]([]int{1, 2, 3, 4, 5}),
			ll[int]([]int{1, 2, 3, 4, 5}).Head,
			true,
		},
		{
			"populated linkedList i at end",
			ll[int]([]int{1, 2, 3, 4, 5}),
			args{4},
			ll[int]([]int{1, 2, 3, 4, 5}),
			ll[int]([]int{1, 2, 3, 4, 5}).Tail,
			true,
		},
		{
			"populated linkedList beginning < i < end",
			ll[int]([]int{1, 2, 3, 4, 5}),
			args{2},
			ll[int]([]int{1, 2, 3, 4, 5}),
			ll[int]([]int{1, 2, 3, 4, 5}).Head.Next.Next,
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotNode, gotBool := tt.existingLinkedList.GetNode(tt.args.i)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("Get() actuall LL = %v, wantLL %v", tt.existingLinkedList, tt.wantLL)
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("Get() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
			if gotBool != tt.wantBool {
				t.Errorf("Get() gotBool = %v, wantBool %v", gotBool, tt.wantBool)
			}
		})
	}
}

func TestLinkedList_IsEmpty(t *testing.T) {
	type testCase[T any] struct {
		name               string
		existingLinkedList *LinkedList[T]
		wantLL             *LinkedList[T]
		want               bool
	}
	tests := []testCase[int]{
		{"empty linkedList", ll[int]([]int{}), ll[int]([]int{}), true},
		{"head value is nil", ll[int](nil), ll[int](nil), false},
		{"populated linkedList", ll[int]([]int{1, 2, 3}), ll[int]([]int{1, 2, 3}), false},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("Get() actuall LL = %v, wantLL %v", tt.existingLinkedList, tt.wantLL)
			}
			if got := tt.existingLinkedList.IsEmpty(); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_Pop(t *testing.T) {
	type testCase[T any] struct {
		name               string
		existingLinkedList *LinkedList[T]
		wantLL             *LinkedList[T]
		wantVal            T
		wantBool           bool
	}
	values := []int{1, 2, 3}
	tests := []testCase[*int]{
		{
			"emtpy linkedList",
			ll[*int]([]*int{}),
			ll[*int]([]*int{}),
			nil,
			false,
		},
		{
			"head value is nil",
			ll[*int]([]*int{nil}),
			ll[*int]([]*int{}),
			nil,
			true,
		},
		{
			"populated linkedList",
			ll[*int]([]*int{&values[0], &values[1], &values[2]}),
			ll[*int]([]*int{&values[1], &values[2]}),
			&values[0],
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotVal, gotBool := tt.existingLinkedList.Pop()
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("Pop() actual LL = %v, wantLL %v", tt.existingLinkedList, tt.wantLL)
			}
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("Pop() gotVal = %v, wantVal %v", gotVal, tt.wantVal)
			}
			if gotBool != tt.wantBool {
				t.Errorf("Pop() gotBool = %v, wantVal %v", gotBool, tt.wantBool)
			}
		})
	}
}

func TestLinkedList_Push(t *testing.T) {
	type args[T any] struct {
		val T
	}
	type testCase[T any] struct {
		name               string
		existingLinkedList *LinkedList[T]
		args               args[T]
		wantLL             *LinkedList[T]
		wantNode           *Node[T]
	}

	intTests := []testCase[int]{
		{
			"int: push int to empty ll",
			ll[int]([]int{}),
			args[int]{1},
			ll[int]([]int{1}),
			ll[int]([]int{1}).Head,
		},
		{
			"int: push int to existing ll",
			ll[int]([]int{1}),
			args[int]{2},
			ll[int]([]int{1, 2}),
			ll[int]([]int{1, 2}).Head.Next,
		},
	}

	emptyListOfLists := make([][]string, 1)
	emptyListOfLists[0] = make([]string, 0)
	listTests := []testCase[[]string]{
		{
			"list: push emtpy list to emtpy ll",
			ll[[]string]([][]string{}),
			args[[]string]{make([]string, 0)},
			ll[[]string](emptyListOfLists),
			ll[[]string](emptyListOfLists).Head,
		},
		{
			"list: push some list to emtpy ll",
			ll[[]string]([][]string{}),
			args[[]string]{[]string{"new-first"}},
			ll[[]string]([][]string{{"new-first"}}),
			ll[[]string]([][]string{{"new-first"}}).Head,
		},
		{
			"list: push empty list to populated ll",
			ll[[]string]([][]string{{"first"}, {"second"}}),
			args[[]string]{make([]string, 0)},
			ll[[]string]([][]string{{"first"}, {"second"}, make([]string, 0)}),
			ll[[]string]([][]string{{"first"}, {"second"}, make([]string, 0)}).Head.Next.Next,
		},
		{
			"list: push some list to populated ll",
			ll[[]string]([][]string{{"first"}, {"second"}, {"third"}}),
			args[[]string]{[]string{"new-fourth"}},
			ll[[]string]([][]string{{"first"}, {"second"}, {"third"}, {"new-fourth"}}),
			ll[[]string]([][]string{{"first"}, {"second"}, {"third"}, {"new-fourth"}}).Head.Next.Next.Next,
		},
	}

	stringPointer := ""
	pointerTests := []testCase[*string]{
		{
			"pointer: push nil to emtpy ll",
			ll[*string]([]*string{}),
			args[*string]{nil},
			ll[*string]([]*string{nil}),
			ll[*string]([]*string{nil}).Head,
		},
		{
			"pointer: push nil to existing ll",
			ll[*string]([]*string{&stringPointer}),
			args[*string]{nil},
			ll[*string]([]*string{&stringPointer, nil}),
			ll[*string]([]*string{&stringPointer, nil}).Head.Next,
		},
	}

	for _, tt := range intTests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			wantNode := node[int](tt.args.val, nil)

			gotNode := tt.existingLinkedList.Push(tt.args.val)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("NewLinkedList() = %v, wantVal %v, wantNode %v", tt.existingLinkedList, tt.wantLL, wantNode)
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("Get() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
		})
	}
	for _, tt := range listTests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			wantNode := node[[]string](tt.args.val, nil)

			gotNode := tt.existingLinkedList.Push(tt.args.val)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("NewLinkedList() = %v, wantVal %v, wantNode %v", tt.existingLinkedList, tt.wantLL, wantNode)
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("Get() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
		})
	}
	for _, tt := range pointerTests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			wantNode := node[*string](tt.args.val, nil)

			gotNode := tt.existingLinkedList.Push(tt.args.val)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("NewLinkedList() = %v, wantVal %v, wantNode %v", tt.existingLinkedList, tt.wantLL, wantNode)
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("Get() gotNode = %v, wantNode %v", gotNode, tt.wantNode)
			}
		})
	}
}

func TestNewLinkedList(t *testing.T) {
	type args[T any] struct {
		elems []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want *LinkedList[T]
	}
	tests := []testCase[int]{
		{"nil", args[int]{[]int{}}, ll[int]([]int{})},
		{"1 item", args[int]{[]int{1}}, ll[int]([]int{1})},
		{"5 items", args[int]{[]int{1, 2, 3, 99, 1001}}, ll[int]([]int{1, 2, 3, 99, 1001})},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			//t.Skip()
			if got := NewLinkedList(tt.args.elems...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, wantVal %v", got, tt.want)
			}
		})
	}
}

func TestLinkedList_PushFront(t *testing.T) {
	type args[T any] struct {
		val T
	}
	type testCase[T any] struct {
		name               string
		existingLinkedList *LinkedList[T]
		args               args[T]
		wantLL             *LinkedList[T]
		wantNode           *Node[T]
	}
	tests := []testCase[int]{
		{
			"empty linkedList",
			ll[int]([]int{}),
			args[int]{1},
			ll[int]([]int{1}),
			node[int](1, nil),
		},
		{
			"populated linkedList",
			ll[int]([]int{1, 2, 3}),
			args[int]{0},
			ll[int]([]int{0, 1, 2, 3}),
			ll[int]([]int{0, 1, 2, 3}).Head,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			if gotNode := tt.existingLinkedList.PushFront(tt.args.val); !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("PushFront() = %v, wantNode %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("Pop() actual LL = %v, wantLL %v", tt.existingLinkedList, tt.wantLL)
			}
		})
	}
}

func TestLinkedList_Insert(t *testing.T) {
	type args[T any] struct {
		val T
		i   uint
	}
	type testCase[T any] struct {
		name               string
		existingLinkedList *LinkedList[T]
		args               args[T]
		wantLL             *LinkedList[T]
		wantNode           *Node[T]
		wantError          error
	}
	tests := []testCase[int]{
		{
			"insert to empty linkedList",
			ll[int]([]int{}),
			args[int]{10, 0},
			ll[int]([]int{10}),
			node[int](10, nil),
			nil,
		},
		{
			"insert to populated LL at i = 0",
			ll[int]([]int{1, 2, 3, 4}),
			args[int]{10, 0},
			ll[int]([]int{10, 1, 2, 3, 4}),
			ll[int]([]int{10, 1, 2, 3, 4}).Head,
			nil,
		},
		{
			"insert to populated LL at 0 < i < ll.Size - 1",
			ll[int]([]int{1, 2, 3, 4}),
			args[int]{10, 2},
			ll[int]([]int{1, 2, 10, 3, 4}),
			ll[int]([]int{1, 2, 10, 3, 4}).Head.Next.Next,
			nil,
		},
		{
			"insert to populated LL at i = ll.Size",
			ll[int]([]int{1, 2, 3, 4}),
			args[int]{10, 4},
			ll[int]([]int{1, 2, 3, 4, 10}),
			ll[int]([]int{1, 2, 3, 4, 10}).Tail,
			nil,
		},
		{
			"insert out of bounds",
			ll[int]([]int{1, 2, 3, 4}),
			args[int]{10, 10},
			ll[int]([]int{1, 2, 3, 4}),
			nil,
			errors.New("index out of range"),
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, tt.name)
			gotNode, gotError := tt.existingLinkedList.Insert(tt.args.val, tt.args.i)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("Pop() actual LL = %v, wantLL %v", tt.existingLinkedList, tt.wantLL)
			}
			if !reflect.DeepEqual(gotNode, tt.wantNode) {
				t.Errorf("Insert() gotNode = %v, want %v", gotNode, tt.wantNode)
			}
			if !reflect.DeepEqual(gotError, tt.wantError) {
				t.Errorf("Insert() gotError = %v, want %v", gotError, tt.wantError)
			}
		})
	}
}
