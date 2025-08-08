package binary_search

import "cmp"

func binarySearch[T cmp.Ordered](arr []T, target T) int {
	left := 0
	right := len(arr) - 1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == target {
			return mid
		}

		if target > arr[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}
