package search

// BinarySearch
//
// comp return:
//
// - equal: 0
//
// - Sorted towards A[0]: negative number
//
// - Sorted towards A[last]: positive number
func BinarySearch[T any](A []T, searchTerm T, comp func(searchTerm, itemB T) int) (index int, found bool) {
	if comp == nil {
		// writing a default comparator for a generic function sucks extremely hard in go, so I just gave up and do this.
		panic("comparator needed")
	}

	b := 0
	e := len(A) - 1
	for b <= e {
		m := (b + e) / 2

		c := comp(searchTerm, A[m])
		if c == 0 {
			return m, true
		} else if c > 0 {
			b = m + 1
		} else {
			e = m - 1
		}
	}

	return index, found
}
