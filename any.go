package gogo

/*
Any ...
Return true if any elements of the iterable are true, the iterable is empty, return false.
*/
func Any(iterable []bool) bool {
	for _, element := range iterable {
		if element {
			return true
		}
	}
	return false
}
