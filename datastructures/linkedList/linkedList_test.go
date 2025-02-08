package linkedList

import (
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
//		want  T
//		want1 bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := tt.l.Get(tt.args.i)
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Get() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("Get() got1 = %v, want %v", got1, tt.want1)
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
//		want bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := tt.l.IsEmpty(); got != tt.want {
//				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
//
//func TestLinkedList_Pop(t *testing.T) {
//	type testCase[T any] struct {
//		name  string
//		l     LinkedList[T]
//		want  T
//		want1 bool
//	}
//	tests := []testCase[ /* TODO: Insert concrete types here */ ]{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			got, got1 := tt.l.Pop()
//			if !reflect.DeepEqual(got, tt.want) {
//				t.Errorf("Pop() got = %v, want %v", got, tt.want)
//			}
//			if got1 != tt.want1 {
//				t.Errorf("Pop() got1 = %v, want %v", got1, tt.want1)
//			}
//		})
//	}
//}

func TestLinkedList_Push(t *testing.T) {
	emptyListOfLists := make([][]string, 1) // Create an outer list with one item
	emptyListOfLists[0] = nil

	type args[T any] struct {
		val T
	}
	type testCase[T any] struct {
		name               string
		existingLinkedList *LinkedList[T]
		args               args[T]
		want               *LinkedList[T]
	}
	intTests := []testCase[int]{
		{"int: push int to empty ll", ll[int]([]int{}), args[int]{1}, ll[int]([]int{1})},
		{"int: push int to existing ll", ll[int]([]int{1}), args[int]{2}, ll[int]([]int{1, 2})},
	}

	listTests := []testCase[[]string]{
		{"list: push emtpy list to emtpy ll",
			ll[[]string]([][]string{}),
			args[[]string]{make([]string, 0)},
			ll[[]string](make([][]string, 1)),
		},
		{"list: push some list to emtpy ll",
			ll[[]string]([][]string{}),
			args[[]string]{[]string{"new-first"}},
			ll[[]string]([][]string{{"new-first"}}),
		},
		{"list: push empty list to existing ll",
			ll[[]string]([][]string{{"first"}, {"second"}}),
			args[[]string]{make([]string, 0)},
			ll[[]string]([][]string{{"first"}, {"second"}}),
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
			ll[*string](nil),
		},
		{"pointer: push nil to existing ll",
			ll[*string]([]*string{&stringPointer}),
			args[*string]{nil},
			ll[*string]([]*string{&stringPointer, nil}),
		},
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.existingLinkedList.Push(tt.args.val)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", tt.existingLinkedList, tt.want)
			}
		})
	}
	for _, tt := range listTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.existingLinkedList.Push(tt.args.val)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", tt.existingLinkedList, tt.want)
			}
		})
	}
	for _, tt := range pointerTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.existingLinkedList.Push(tt.args.val)
			if !reflect.DeepEqual(tt.existingLinkedList, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", tt.existingLinkedList, tt.want)
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
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedList(tt.args.elems...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ll[T any](vals []T) *LinkedList[T] {
	if len(vals) == 0 {
		return &LinkedList[T]{Head: nil, Tail: nil, size: 0}
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
