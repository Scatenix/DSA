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
// - Sort a towards A[0]: negative number
//
// - Sort a towards A[n]: positive number
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
// - Sort a towards A[0]: negative number
//
// - Sort a towards A[n]: positive number
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
// Runtime: Runtime: Ω(n), O(n²)
//
// comp - comparator should return
//
// - Equal: 0
//
// - Sort a towards A[0]: negative number
//
// - Sort a towards A[n]: positive number
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

func MergeSort[T any](A []T, comp func(a, b T) int) (sorted []T) {
	return A
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
