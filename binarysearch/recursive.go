package binarysearch

// Recursive binary search
func Recursive(array []int, target int) int {
	return recursiveSearch(array, target, 0, len(array)-1)
}

func recursiveSearch(array []int, target int, l int, h int) int {
	if l > h {
		return -1
	}

	m := l + (h-l)/2

	if array[m] == target {
		return m
	} else if array[m] < target {
		return recursiveSearch(array, target, m+1, h)
	} else {
		return recursiveSearch(array, target, l, m-1)
	}
}
