package search

import (
	"dsa/util/sugar"
	"reflect"
	"testing"
)

func TestBinarySearchInt(t *testing.T) {
	type args[T any] struct {
		A          []T
		searchTerm T
		comp       func(A T, B T) bool
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantArray []T
		wantVal   T
		wantFound bool
	}
	tests := []testCase[int]{
		{
			"empty array",
			args[int]{
				[]int{},
				1,
				nil,
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
				nil,
			},
			[]int{2},
			2,
			true,
		},
		{
			"single array, odd, not found ",
			args[int]{
				[]int{2},
				1,
				nil,
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
				nil,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			16,
			true,
		},
		{
			"big array, even, not found, greater than",
			args[int]{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				21,
				nil,
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
				nil,
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
				nil,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21},
			0,
			false,
		},
		{
			"big array, duplicate values, found",
			args[int]{
				[]int{1, 2, 3, 4, 5, 6, 7, 8, 8, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
				8,
				nil,
			},
			[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20},
			8,
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, t.Name())
			gotVal, gotFound := BinarySearch(tt.args.A, tt.args.searchTerm, tt.args.comp)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("int BinarySearch() = %v, want %v", gotVal, tt.wantVal)
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
		comp       func(A T, B T) bool
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantArray []T
		wantVal   T
		wantFound bool
	}
	tests := []testCase[string]{
		{
			"empty array",
			args[string]{
				[]string{},
				"1",
				nil,
			},
			[]string{},
			"0",
			false,
		},
		{
			"single array, odd, found",
			args[string]{
				[]string{"2"},
				"2",
				nil,
			},
			[]string{"2"},
			"2",
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, t.Name())
			gotVal, gotFound := BinarySearch(tt.args.A, tt.args.searchTerm, tt.args.comp)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("string BinarySearch() = %v, want %v", gotVal, tt.wantVal)
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
		comp       func(A T, B T) bool
	}
	type testCase[T any] struct {
		name      string
		args      args[T]
		wantArray []T
		wantVal   T
		wantFound bool
	}
	tests := []testCase[complexType]{
		{
			"empty array",
			args[complexType]{
				[]complexType{},
				complexType{0, ""},
				nil,
			},
			[]complexType{},
			complexType{0, "a"},
			false,
		},
		{
			"single array, odd, found",
			args[complexType]{
				[]complexType{complexType{0, ""}},
				complexType{0, ""},
				nil,
			},
			[]complexType{complexType{0, ""}},
			complexType{0, ""},
			true,
		},
	}
	for _, tt := range tests {
		println()
		t.Run(tt.name, func(t *testing.T) {
			defer sugar.Lite(t, t.Name())
			gotVal, gotFound := BinarySearch(tt.args.A, tt.args.searchTerm, tt.args.comp)
			if !reflect.DeepEqual(gotVal, tt.wantVal) {
				t.Errorf("complexType BinarySearch() = %v, want %v", gotVal, tt.wantVal)
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
