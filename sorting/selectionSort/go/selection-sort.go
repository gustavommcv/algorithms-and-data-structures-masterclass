package selectionSort

import "cmp"

func SelectionSort[T cmp.Ordered](arr []T) {
	for i := range arr {
		lowestIndex := i

		for j := i + 1; j < len(arr); j++ {
			if arr[lowestIndex] > arr[j] {
				lowestIndex = j
			}
		}

		if arr[lowestIndex] != arr[i] {
			arr[i], arr[lowestIndex] = arr[lowestIndex], arr[i]
		}
	}
}
