package search

func BinarySearch[T any](A []T, searchTerm T, comp func(searchTerm, itemB T) int) (index int, found bool) {
	if comp == nil {
		// writing a default comparator for a generic function sucks extremely hard in go, so I just gave up and do this.
		panic("comparator needed")
	}

	return index, found
}
