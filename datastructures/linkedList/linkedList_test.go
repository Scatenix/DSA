package linkedList

import (
	"dsa/util/sugar"
	"reflect"
	"testing"
)

//func TestLinkedList_Get(t *testing.T) {
//	type args struct {
//		i int
//	}
//	type testCase[T any] struct {
//		name  string
//		l     LinkedList[T]
//		args  args
//		wantLL  T
//		want1 bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := tt.l.Get(tt.args.i)
//			if !reflect.DeepEqual(got, tt.wantLL) {
//				t.Errorf("Get() got = %v, wantLL %v", got, tt.wantLL)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("Get() got1 = %v, wantLL %v", got1, tt.want1)
//			}
//		})
//	}
//}
//
//func TestLinkedList_Insert(t *testing.T) {
//	type args[T any] struct {
//		val T
//		i   int
//	}
//	type testCase[T any] struct {
//		name string
//		l    LinkedList[T]
//		args args[T]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.l.Insert(tt.args.val, tt.args.i)
//		})
//	}
//}
//
//func TestLinkedList_IsEmpty(t *testing.T) {
//	type testCase[T any] struct {
//		name string
//		l    LinkedList[T]
//		wantLL bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.l.IsEmpty(); got != tt.wantLL {
//				t.Errorf("IsEmpty() = %v, wantLL %v", got, tt.wantLL)
//			}
//		})
//	}
//}
//
//func TestLinkedList_Pop(t *testing.T) {
//	type testCase[T any] struct {
//		name  string
//		l     LinkedList[T]
//		wantLL  T
//		want1 bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := tt.l.Pop()
//			if !reflect.DeepEqual(got, tt.wantLL) {
//				t.Errorf("Pop() got = %v, wantLL %v", got, tt.wantLL)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("Pop() got1 = %v, wantLL %v", got1, tt.want1)
//			}
//		})
//	}
//}

func TestLinkedList_Push(t *testing.T) {
	type args[T any] struct {
		val T
	}
	type testCase[T any] struct {
		name               string
		existingLinkedList *LinkedList[T]
		args               args[T]
		wantLL             *LinkedList[T]
	}

	intTests := []testCase[int]{
		{"int: push int to empty ll", ll[int]([]int{}), args[int]{1}, ll[int]([]int{1})},
		{"int: push int to existing ll", ll[int]([]int{1}), args[int]{2}, ll[int]([]int{1, 2})},
	}

	emptyListOfLists := make([][]string, 1)
	emptyListOfLists[0] = make([]string, 0)
	listTests := []testCase[[]string]{
		{"list: push emtpy list to emtpy ll",
			ll[[]string]([][]string{}),
			args[[]string]{make([]string, 0)},
			ll[[]string](emptyListOfLists),
		},
		{"list: push some list to emtpy ll",
			ll[[]string]([][]string{}),
			args[[]string]{[]string{"new-first"}},
			ll[[]string]([][]string{{"new-first"}}),
		},
		{"list: push empty list to existing ll",
			ll[[]string]([][]string{{"first"}, {"second"}}),
			args[[]string]{make([]string, 0)},
			ll[[]string]([][]string{{"first"}, {"second"}, make([]string, 0)}),
		},
		{"list: push some list to existing ll",
			ll[[]string]([][]string{{"first"}, {"second"}, {"third"}}),
			args[[]string]{[]string{"new-fourth"}},
			ll[[]string]([][]string{{"first"}, {"second"}, {"third"}, {"new-fourth"}}),
		},
	}

	stringPointer := ""
	pointerTests := []testCase[*string]{
		{"pointer: push nil to emtpy ll",
			ll[*string]([]*string{}),
			args[*string]{nil},
			ll[*string]([]*string{nil}),
		},
		{"pointer: push nil to existing ll",
			ll[*string]([]*string{&stringPointer}),
			args[*string]{nil},
			ll[*string]([]*string{&stringPointer, nil}),
		},
	}

	for _, tt := range intTests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(tt.name, t)
			wantNode := node[int](tt.args.val, nil)

			tt.existingLinkedList.Push(tt.args.val)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("NewLinkedList() = %v, wantLL %v, wantNode %v", tt.existingLinkedList, tt.wantLL, wantNode)
			}
		})
	}
	for _, tt := range listTests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(tt.name, t)
			wantNode := node[[]string](tt.args.val, nil)

			tt.existingLinkedList.Push(tt.args.val)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("NewLinkedList() = %v, wantLL %v, wantNode %v", tt.existingLinkedList, tt.wantLL, wantNode)
			}
		})
	}
	for _, tt := range pointerTests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(tt.name, t)
			wantNode := node[*string](tt.args.val, nil)

			tt.existingLinkedList.Push(tt.args.val)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.wantLL) {
				t.Errorf("NewLinkedList() = %v, wantLL %v, wantNode %v", tt.existingLinkedList, tt.wantLL, wantNode)
			}
		})
	}
}

//func TestLinkedList_PushFront(t *testing.T) {
//	type args[T any] struct {
//		val T
//	}
//	type testCase[T any] struct {
//		name string
//		l    LinkedList[T]
//		args args[T]
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.l.PushFront(tt.args.val)
//		})
//	}
//}
//
//func TestLinkedList_Remove(t *testing.T) {
//	type args struct {
//		i int
//	}
//	type testCase[T any] struct {
//		name string
//		l    LinkedList[T]
//		args args
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			tt.l.Remove(tt.args.i)
//		})
//	}
//}

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
			defer sugar.Lite(tt.name, t)
			//t.Skip()
			if got := NewLinkedList(tt.args.elems...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, wantLL %v", got, tt.want)
			}
		})
	}
}

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
	ll := &LinkedList[T]{node, nil, len(vals)}

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
