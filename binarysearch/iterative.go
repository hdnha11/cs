package binarysearch

// Iterative binary search
func Iterative(array []int, target int) int {
	l := 0
	h := len(array) - 1

	for l <= h {
		m := l + (h-l)/2

		if array[m] == target {
			return m
		} else if array[m] < target {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return -1
}
