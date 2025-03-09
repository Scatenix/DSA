package arraySort

// BubbleSort sorts an array A and returns it as sorted.
// Sorting operation is done directly on the instance of array A.
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
		for j := 0; j < i-1; j++ {
			c := comp(A[j], A[j+1])
			if c > 0 {
				//sawp
				tmp := A[j]
				A[j] = A[j+1]
				A[j+1] = tmp
			}
		}
	}
	return A
}
