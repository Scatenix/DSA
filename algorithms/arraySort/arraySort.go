package arraySort

// BubbleSort sorts an array A and returns it as sorted.
// Sorting operation is done directly on the instance of array A.
//
// Runtime: Ω(n), O(N²) (Hint: Omega(n) because of the 'swapped' variable)
//
// comp - comparator should return
//
// - Equal: 0
//
// - Sorted towards A[0]: negative number
//
// - Sorted towards A[last]: positive number
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
// Runtime: Theta(N²)
//
// comp - comparator should return
//
// - Equal: 0
//
// - Sorted towards A[0]: negative number
//
// - Sorted towards A[last]: positive number
func SelectionSort[T any](A []T, comp func(a, b T) int) (sorted []T) {
	if comp == nil {
		panic("Provided comparator was nil")
	}

	for i := 0; i < len(A)-1; i++ {
		smallest := i
		for j := i + 1; j < len(A)-1; j++ {
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
