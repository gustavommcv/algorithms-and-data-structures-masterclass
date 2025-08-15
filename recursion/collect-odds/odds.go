package main

func main() {}

func collectOddValues(arr []int) []int {
	var oddNumbers []int

	var helper func(i []int)
	helper = func(helperInput []int) {
		if len(helperInput) == 0 {
			return
		}

		if helperInput[0]%2 == 1 {
			oddNumbers = append(oddNumbers, helperInput[0])
		}

		helper(helperInput[1:])
	}

	helper(arr)

	return oddNumbers
}
