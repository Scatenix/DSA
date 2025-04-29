package arraySort

// BubbleSort sorts an array A and returns it as sorted.
// Sorting operation is done directly on the instance of array A.
//
// Runtime: Ω(n), O(n²) (Hint: Ω(n) because of the 'swapped' variable)
//
// comp - comparator should return
//
// - Equal: 0
//
// - Sort 'a' towards A[0], 'a' before 'b': negative number
//
// - Sort 'a' towards A[n], 'a' after 'b': positive number
func BubbleSort[T any](A []T, comp func(a, b T) int) (sorted []T) {
	if comp == nil {
		panic("Provided comparator was nil")
	}

	for i := len(A); i > 0; i-- {

		// This swapped variable allows us an early exit if the array is already sorted before we go through both for loops.
		swapped := false
		for j := 0; j < i-1; j++ {
			c := comp(A[j], A[j+1])
			if c > 0 {
				// sawp
				tmp := A[j]
				A[j] = A[j+1]
				A[j+1] = tmp
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return A
}

// SelectionSort sorts an array A and returns it as sorted.
// Sorting operation is done directly on the instance of array A.
//
// Runtime: Theta(n²)
//
// comp - comparator should return
//
// - Equal: 0
//
// - Sort 'a' towards A[0], 'a' before 'b': negative number
//
// - Sort 'a' towards A[n], 'a' after 'b': positive number
func SelectionSort[T any](A []T, comp func(a, b T) int) (sorted []T) {
	if comp == nil {
		panic("Provided comparator was nil")
	}

	for i := 0; i < len(A)-1; i++ {
		smallest := i
		for j := i + 1; j < len(A); j++ {
			if comp(A[j], A[smallest]) < 0 {
				smallest = j
			}
		}
		// swap
		tmp := A[i]
		A[i] = A[smallest]
		A[smallest] = tmp
	}

	return A
}

// InsertionSort sorts an array A and returns it as sorted.
// Sorting operation is done directly on the instance of array A.
//
// Runtime: Ω(n), O(n²)
//
// comp - comparator should return
//
// - Equal: 0
//
// - Sort 'a' towards A[0], 'a' before 'b': negative number
//
// - Sort 'a' towards A[n], 'a' after 'b': positive number
func InsertionSort[T any](A []T, comp func(a, b T) int) (sorted []T) {
	if comp == nil {
		panic("Provided comparator was nil")
	}

	for i := 1; i < len(A); i++ {
		key := A[i]
		j := i - 1
		for j >= 0 && comp(A[j], key) > 0 {
			A[j+1] = A[j]
			j--
		}
		A[j+1] = key
	}

	return A
}

// MergeSortInt sorts an array A of type int and returns it as sorted.
// Sorting operation is done on a new array. Space complexity should be 2n (peak memory usage) -> Theta(n)
// Why not 2n log n? because
//
// Using pure functions. A lot simpler implementation than using impure functions in my opinion.
// Combined with the fact that it is not generic, this should be the easiest possible implementation for merge sort.
//
// Runtime: Theta(n log n)
//
// comp - comparator should return
//
// - Equal: 0
//
// - Sort 'a' towards A[0], 'a' before 'b': negative number
//
// - Sort 'a' towards A[n], 'a' after 'b': positive number
func MergeSortInt(A []int) []int {
	if len(A) <= 1 {
		return A
	}
	q := len(A) / 2

	B := MergeSortInt(A[:q]) // Remember, it is [inclusive:exclusive]
	C := MergeSortInt(A[q:]) // That's why this is not [q+1:]!!! We did not take q into account yet.
	return mergeSortIntMerge(B, C)
}

func mergeSortIntMerge(B, C []int) (A []int) {
	A = make([]int, len(B)+len(C))

	// indices for A, B and C arrays. Very easy to read imo, even though this is not following naming conventions.
	a := 0
	b := 0
	c := 0

	for b < len(B) && c < len(C) {
		if B[b] <= C[c] {
			A[a] = B[b]
			b++
		} else {
			A[a] = C[c]
			c++
		}
		a++
	}

	for ; b < len(B); b++ {
		A[a] = B[b]
		a++
	}
	for ; c < len(C); c++ {
		A[a] = C[c]
		a++
	}
	return A
}

// MergeSort sorts an array A of generic type and returns it as sorted.
// Sorting operation is done on a new array. Space complexity should be 2n (peak memory usage) -> Theta(n)
//
// Using impure functions. A lot more complicated implementation than using pure functions in my opinion.
//
// Runtime: Theta(n log n)
//
// comp - comparator should return
//
// - Equal: 0
//
// - Sort 'a' towards A[0], 'a' before 'b': negative number
//
// - Sort 'a' towards A[n], 'a' after 'b': positive number
func MergeSort[T any](A []T, comp func(a, b T) int) (sorted []T) {
	if comp == nil {
		panic("Provided comparator was nil")
	}

	mergeSortRecurse(A, comp, 0, len(A))
	return A
}

func mergeSortRecurse[T any](A []T, comp func(a, b T) int, p, r int) {
	if p >= r {
		return
	}
	q := (p + r) / 2

	mergeSortRecurse[T](A, comp, p, q)
	mergeSortRecurse[T](A, comp, q+1, r)
	mergeSortMerge[T](A, comp, p, q, r)
}

func mergeSortMerge[T any](A []T, comp func(a, b T) int, p, q, r int) {
	// This whole stuff with copying arrays to B and C could have been completely avoided if this function takes in the 2 arrays to be merged.
	// This then means that the mergeSortRecurse function needs to return the array produced by this mergeSortMerge function which then in turn get
	// passed to this merge function again on the next round as one of the 2 parameters.
	n1 := q - p + 1
	n2 := r - q

	B := make([]T, n1)
	copy(B, A[p:q+1])

	var C []T
	if r < len(A) {
		C = make([]T, n2)
		copy(C, A[q+1:r+1])
	} else if q+1 < len(A) {
		C = make([]T, n2-1)
		copy(C, A[q+1:])
	}

	// indices for A, B and C arrays. Very easy to read imo, even though this is not following naming conventions.
	a := p
	b := 0
	c := 0
	for b < n1 && c < n2 {
		// This whole if else else block would be much simpler if it wasn't for the genericnes of the function.
		res := 0
		if b < len(B) && c < len(C) {
			res = comp(B[b], C[c])
		} else if b >= len(B) {
			res = 1
		} else if c >= len(C) {
			res = -1
		}

		if res <= 0 {
			A[a] = B[b]
			b++
		} else {
			A[a] = C[c]
			c++
		}
		a++
	}

	// next to loops are also a tad easier if I just used the 2 arrays as function parameters instead of the impure function mess.
	// Copy remaining elements after one of B or C has been emptied
	for b < n1 && a < len(A) {
		A[a] = B[b]
		b++
		a++
	}
	for c < n2 && C != nil && a < len(A) {
		A[a] = C[c]
		c++
		a++
	}
}

func QuickSort[T any](A []T, comp func(a, b T) int) (sorted []T) {
	return A
}

// Runtime: Theta(n log n)
func HeapSort() {}

// Runtime: Theta(n + k)
// Works only on Integers
func CoutingSort() {}

// Runtime: Ω(n log n), O(n²)
func ShellSort() {}

// Runtime: Theta(n log n)
func TimSort() {}

// Runtime: Theta(d * (n + b))
// d: Amount of numbers in the largest number
// n: number of elements
// b: base of the numbers (like binary base 2, decimal base 10)
func RadixSort() {}
