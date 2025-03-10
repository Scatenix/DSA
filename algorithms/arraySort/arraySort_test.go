package arraySort

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

func TestSortAlgorithmsPanic(t *testing.T) {
	type args struct {
		A    []int
		comp func(a, b int) int
	}
	type testCase struct {
		name       string
		args       args
		wantSorted []int
	}
	type functions struct {
		name string
		fn   func(A []int, comp func(a, b int) int) (sorted []int)
	}
	tests := []testCase{
		{
			"no comparator, panic",
			args{
				[]int{1},
				nil,
			},
			nil,
		},
	}
	funcs := []functions{
		// This was the best way I found to test different sorting algorithms with one test function.
		// Would be a bit less verbose if the functions were not generic.
		{"BubbleSort", func(A []int, comp func(a, b int) int) []int { return BubbleSort(A, comp) }},
		{"SelectionSort", func(A []int, comp func(a, b int) int) []int { return SelectionSort(A, comp) }},
		//{"MergeSort", func(A []int, comp func(a, b int) int) []int { return MergeSort(A, comp) }},
		//{"QuickSort", func(A []int, comp func(a, b int) int) []int { return QuickSort(A, comp) }},
	}
	for _, fn := range funcs {
		for _, tt := range tests {
			println()
			// Hint: cannot put 'fn.name+": "+tt.name' as name for t.Run, because GoLand then cannot match the
			// individual testcases anymore. Meaning the run/success/failed icon next to the testcase will be missing completely.
			t.Run(tt.name, func(t *testing.T) {
				defer sugar.Lite(t, fn.name+": "+tt.name)
				defer func() {
					if r := recover(); r == nil {
						t.Errorf("BinarySearch(): Expected to panic, but it did not.")
					}
				}()

				fn.fn(tt.args.A, tt.args.comp)
			})
		}
	}
}

func TestSortAlgorithmsInt(t *testing.T) {
	type args struct {
		A    []int
		comp func(a, b int) int
	}
	type testCase struct {
		name       string
		args       args
		wantSorted []int
	}
	type functions struct {
		name string
		fn   func(A []int, comp func(a, b int) int) (sorted []int)
	}
	tests := []testCase{
		{
			"empty list",
			args{
				[]int{},
				intComp,
			},
			[]int{},
		},
		{
			"single array",
			args{
				[]int{1},
				intComp,
			},
			[]int{1},
		},
		{
			"big array",
			args{
				[]int{7, 2, 5, 1, 8, 0, 3, 6, 4, 9},
				intComp,
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"duplicates array",
			args{
				[]int{7, 2, 5, 0, 1, 8, 3, 3, 3, 6, 5, 4, 9},
				intComp,
			},
			[]int{0, 1, 2, 3, 3, 3, 4, 5, 5, 6, 7, 8, 9},
		},
		{
			"reverse comparator",
			args{
				[]int{7, 2, 0, 1, 8, 3, 6, 5, 4, 9},
				func(a int, b int) int {
					return (a - b) * -1
				},
			},
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
		{
			"already sorted",
			args{
				[]int{0, 1, 2, 3, 3, 3, 4, 5, 5, 6, 7, 8, 9},
				intComp,
			},
			[]int{0, 1, 2, 3, 3, 3, 4, 5, 5, 6, 7, 8, 9},
		},
	}
	funcs := []functions{
		// This was the best way I found to test different sorting algorithms with one test function.
		// Would be a bit less verbose if the functions were not generic.
		{"BubbleSort", func(A []int, comp func(a, b int) int) []int { return BubbleSort(A, comp) }},
		{"SelectionSort", func(A []int, comp func(a, b int) int) []int { return SelectionSort(A, comp) }},
		//{"MergeSort", func(A []int, comp func(a, b int) int) []int { return MergeSort(A, comp) }},
		//{"QuickSort", func(A []int, comp func(a, b int) int) []int { return QuickSort(A, comp) }},
	}
	for _, fn := range funcs {
		for _, tt := range tests {
			println()
			// Hint: cannot put 'fn.name+": "+tt.name' as name for t.Run, because GoLand then cannot match the
			// individual testcases anymore. Meaning the run/success/failed icon next to the testcase will be missing completely.
			t.Run(tt.name, func(t *testing.T) {
				defer sugar.Lite(t, fn.name+": "+tt.name)
				if gotSorted := fn.fn(tt.args.A, tt.args.comp); !reflect.DeepEqual(gotSorted, tt.wantSorted) {
					t.Errorf("BubbleSort() = %v, want %v", gotSorted, tt.wantSorted)
				}
			})
		}
	}
}

func TestSortAlgorithmsString(t *testing.T) {
	type args struct {
		A    []string
		comp func(a, b string) int
	}
	type testCase struct {
		name       string
		args       args
		wantSorted []string
	}
	type functions struct {
		name string
		fn   func(A []string, comp func(a, b string) int) (sorted []string)
	}
	tests := []testCase{
		{
			"empty list",
			args{
				[]string{},
				stringComp,
			},
			[]string{},
		},
		{
			"single array",
			args{
				[]string{"abc"},
				stringComp,
			},
			[]string{"abc"},
		},
		{
			"big array",
			args{
				[]string{"vwx", "jkl", "abc", "mno", "zcc", "stu", "zee", "ghi", "yza", "pqr", "zbb", "zdd", "zaa", "def"},
				stringComp,
			},
			[]string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "vwx", "yza", "zaa", "zbb", "zcc", "zdd", "zee"},
		},
		{
			"duplicates array",
			args{
				[]string{"vwx", "jkl", "abc", "mno", "zcc", "zbb", "stu", "zee", "stu", "ghi", "yza", "pqr", "zbb", "zdd", "stu", "zaa", "def"},
				stringComp,
			},
			[]string{"abc", "def", "ghi", "jkl", "mno", "pqr", "stu", "stu", "stu", "vwx", "yza", "zaa", "zbb", "zbb", "zcc", "zdd", "zee"},
		},
	}
	funcs := []functions{
		// This was the best way I found to test different sorting algorithms with one test function.
		// Would be a bit less verbose if the functions were not generic.
		{"BubbleSort", func(A []string, comp func(a, b string) int) []string { return BubbleSort(A, comp) }},
		{"SelectionSort", func(A []string, comp func(a, b string) int) []string { return SelectionSort(A, comp) }},
		//{"MergeSort", func(A []string, comp func(a, b string) int) []string { return MergeSort(A, comp) }},
		//{"QuickSort", func(A []string, comp func(a, b string) int) []string { return QuickSort(A, comp) }},
	}
	for _, fn := range funcs {
		for _, tt := range tests {
			println()
			// Hint: cannot put 'fn.name+": "+tt.name' as name for t.Run, because GoLand then cannot match the
			// individual testcases anymore. Meaning the run/success/failed icon next to the testcase will be missing completely.
			t.Run(tt.name, func(t *testing.T) {
				defer sugar.Lite(t, fn.name+": "+tt.name)
				if gotSorted := fn.fn(tt.args.A, tt.args.comp); !reflect.DeepEqual(gotSorted, tt.wantSorted) {
					t.Errorf("BubbleSort() = %v, want %v", gotSorted, tt.wantSorted)
				}
			})
		}
	}
}
