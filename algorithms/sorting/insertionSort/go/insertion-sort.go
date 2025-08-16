package insertionSort

import "cmp"

func InsertionSort[T cmp.Ordered](arr []T) {
	for i := 1; i < len(arr); i++ {
		current := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > current {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = current
	}
}
