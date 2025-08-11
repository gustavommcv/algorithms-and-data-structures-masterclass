package bubbleSort

import "cmp"

func BubbleSort[T cmp.Ordered](arr []T) {

	for i := len(arr) - 1; i >= 0; i-- {
		noSwaps := true

		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				noSwaps = false
			}
		}

		if noSwaps {
			break
		}
	}

}
