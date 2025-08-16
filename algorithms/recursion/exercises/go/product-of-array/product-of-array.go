package product_of_array

func ProductOfArray(arr []int) int {
	if len(arr) < 1 {
		return 1
	}

	lastElement, newArr := pop(arr)
	return lastElement * ProductOfArray(newArr)
}

func pop(arr []int) (int, []int) {
	return arr[len(arr)-1], arr[:len(arr)-1]
}
