package search

import (
	"dsa/util/sugar"
	"reflect"
	"strings"
	"testing"
)

var intComp = func(a, b int) int {
	return a - b
}

var stringComp = func(a, b string) int {
	return strings.Compare(a, b)
}

var complexComp = func(searchTerm complexType, B complexType) int {
	return strings.Compare(searchTerm.Y, B.Y)
}

func TestBinarySearchPanic(t *testing.T) {
	type args[T any] struct {
		A          []T
		searchTerm T
		comp       func(A T, B T) int
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantArray []T
		wantIndex int
		wantFound bool
	}
	tests := []testCase[int]{
		{
			"no comparator, panic",
			args[int]{
				[]int{1},
				1,
				nil,
			},
			[]int{1},
			0,
			false,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, t.Name())
			defer func() {
				if r := recover(); r == nil {
					t.Errorf("BinarySearch(): Expected to panic, but it did not.")
				}
			}()

			BinarySearch(tt.args.A, tt.args.searchTerm, tt.args.comp)
		})
	}
}

func TestBinarySearchInt(t *testing.T) {
	type args[T any] struct {
		A          []T
		searchTerm T
		comp       func(A T, B T) int
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantArray []T
		wantIndex int
		wantFound bool
	}
	tests := []testCase[int]{
		{
			"empty array",
			args[int]{
				[]int{},
				1,
				intComp,
			},
			[]int{},
			0,
			false,
		},
		{
			"single array, odd, found",
			args[int]{
				[]int{2},
				2,
				intComp,
			},
			[]int{2},
			0,
			true,
		},
		{
			"single array, odd, not found ",
			args[int]{
				[]int{2},
				1,
				intComp,
			},
			[]int{2},
			0,
			false,
		},
		{
			"big array, even, found",
			args[int]{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				16,
				intComp,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			15,
			true,
		},
		{
			"big array, even, not found, greater than",
			args[int]{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				21,
				intComp,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			0,
			false,
		},
		{
			"big array, even, not found, less than",
			args[int]{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				0,
				intComp,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			0,
			false,
		},
		{
			"big array, even, not found, between values",
			args[int]{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
				10,
				intComp,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
			0,
			false,
		},
		{
			"big array, duplicate values, found",
			args[int]{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 8, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				8,
				intComp,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 8, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			7,
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, t.Name())
			gotIndex, gotFound := BinarySearch(tt.args.A, tt.args.searchTerm, tt.args.comp)
			if !reflect.DeepEqual(gotIndex, tt.wantIndex) {
				t.Errorf("int BinarySearch() = %v, want %v", gotIndex, tt.wantIndex)
			}
			if !reflect.DeepEqual(gotFound, tt.wantFound) {
				t.Errorf("int BinarySearch() = %v, want %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(tt.args.A, tt.wantArray) {
				t.Errorf("int BinarySearch() - input array was changed!")
			}
		})
	}
}

func TestBinarySearchString(t *testing.T) {
	type args[T any] struct {
		A          []T
		searchTerm T
		comp       func(A T, B T) int
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantArray []T
		wantIndex int
		wantFound bool
	}
	tests := []testCase[string]{
		{
			"string in single array, found",
			args[string]{
				[]string{"abc"},
				"abc",
				stringComp,
			},
			[]string{"abc"},
			0,
			true,
		},
		{
			"string in single array, not found",
			args[string]{
				[]string{"abc"},
				"def",
				stringComp,
			},
			[]string{"abc"},
			0,
			false,
		},
		{
			"string in big array, found",
			args[string]{
				[]string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yza", "zaa", "zbb", "zcc", "zdd", "zee"},
				"zaa",
				stringComp,
			},
			[]string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yza", "zaa", "zbb", "zcc", "zdd", "zee"},
			9,
			true,
		},
		{
			"string in big array, not found",
			args[string]{
				[]string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yza", "zaa", "zbb", "zcc", "zdd", "zee"},
				"lmn",
				stringComp,
			},
			[]string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yza", "zaa", "zbb", "zcc", "zdd", "zee"},
			0,
			false,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, t.Name())
			gotIndex, gotFound := BinarySearch(tt.args.A, tt.args.searchTerm, tt.args.comp)
			if !reflect.DeepEqual(gotIndex, tt.wantIndex) {
				t.Errorf("string BinarySearch() = %v, want %v", gotIndex, tt.wantIndex)
			}
			if !reflect.DeepEqual(gotFound, tt.wantFound) {
				t.Errorf("string BinarySearch() = %v, want %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(tt.args.A, tt.wantArray) {
				t.Errorf("string BinarySearch() - input array was changed!")
			}
		})
	}
}

type complexType struct {
	X int
	Y string
}

func TestBinarySearchComplex(t *testing.T) {
	type args[T any] struct {
		A          []T
		searchTerm T
		comp       func(A T, B T) int
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantArray []T
		wantIndex int
		wantFound bool
	}
	tests := []testCase[complexType]{
		{
			"complexType in single array, found",
			args[complexType]{
				[]complexType{{1, "abc"}},
				complexType{1, "abc"},
				complexComp,
			},
			[]complexType{{1, "abc"}},
			0,
			true,
		},
		{
			"complexType in single array, not found",
			args[complexType]{
				[]complexType{{1, "abc"}},
				complexType{1, "aaa"},
				complexComp,
			},
			[]complexType{{1, "abc"}},
			0,
			false,
		},
		{
			"complexType in big array, found",
			args[complexType]{
				[]complexType{{1, "abc"}, {2, "def"}, {3, "ghi"}, {4, "jkl"}, {5, "mno"},
					{6, "pqr"}, {7, "stu"}, {8, "vwx"}, {9, "yza"}, {10, "zaa"},
					{11, "zbb"}, {12, "zcc"}, {13, "zdd"}, {14, "zee"}},
				complexType{3, "ghi"},
				complexComp,
			},
			[]complexType{{1, "abc"}, {2, "def"}, {3, "ghi"}, {4, "jkl"}, {5, "mno"},
				{6, "pqr"}, {7, "stu"}, {8, "vwx"}, {9, "yza"}, {10, "zaa"},
				{11, "zbb"}, {12, "zcc"}, {13, "zdd"}, {14, "zee"}},
			2,
			true,
		},
		{
			"complexType in big array, not found",
			args[complexType]{
				[]complexType{{1, "abc"}, {2, "def"}, {3, "ghi"}, {4, "jkl"}, {5, "mno"},
					{6, "pqr"}, {7, "stu"}, {8, "vwx"}, {9, "yza"}, {10, "zaa"},
					{11, "zbb"}, {12, "zcc"}, {13, "zdd"}, {14, "zee"}},
				complexType{8, "lmn"},
				complexComp,
			},
			[]complexType{{1, "abc"}, {2, "def"}, {3, "ghi"}, {4, "jkl"}, {5, "mno"},
				{6, "pqr"}, {7, "stu"}, {8, "vwx"}, {9, "yza"}, {10, "zaa"},
				{11, "zbb"}, {12, "zcc"}, {13, "zdd"}, {14, "zee"}},
			0,
			false,
		},
		{
			"complexType in big array, bad comparator function",
			args[complexType]{
				[]complexType{{1, "abc"}, {2, "def"}, {3, "ghi"}, {4, "jkl"}, {5, "mno"},
					{6, "pqr"}, {7, "stu"}, {8, "vwx"}, {9, "yza"}, {10, "zaa"},
					{11, "zbb"}, {12, "zcc"}, {13, "zdd"}, {14, "zee"}},
				complexType{8, "lmn"},
				func(searchTerm complexType, B complexType) int {
					return strings.Compare(searchTerm.Y, B.Y) * -1
				},
			},
			[]complexType{{1, "abc"}, {2, "def"}, {3, "ghi"}, {4, "jkl"}, {5, "mno"},
				{6, "pqr"}, {7, "stu"}, {8, "vwx"}, {9, "yza"}, {10, "zaa"},
				{11, "zbb"}, {12, "zcc"}, {13, "zdd"}, {14, "zee"}},
			0,
			false,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, t.Name())
			gotIndex, gotFound := BinarySearch(tt.args.A, tt.args.searchTerm, tt.args.comp)
			if !reflect.DeepEqual(gotIndex, tt.wantIndex) {
				t.Errorf("complexType BinarySearch() = %v, want %v", gotIndex, tt.wantIndex)
			}
			if !reflect.DeepEqual(gotFound, tt.wantFound) {
				t.Errorf("complexType BinarySearch() = %v, want %v", gotFound, tt.wantFound)
			}
			if !reflect.DeepEqual(tt.args.A, tt.wantArray) {
				t.Errorf("complexType BinarySearch() - input array was changed!")
			}
		})
	}
}
