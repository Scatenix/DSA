package search

// TODO: if comp = nil, check if T is type string or rune and compare alphabetically. if numeric: standard numerically,
//
//	else panic, or return error
func BinarySearch[T any](A []T, searchTerm T, comp func(A T, B T) bool) (val T, found bool) {
	return val, found
}
