package fib

func fib(input int) int {
	if input == 0 {
		return 0
	}

	fibSequence := make([]int, 1, input)
	fibSequence = append(fibSequence, 1)

	var helperFunc func()
	helperFunc = func() {
		if input == len(fibSequence)-1 {
			return
		}

		fibSequence = append(fibSequence, fibSequence[len(fibSequence)-1]+fibSequence[len(fibSequence)-2])
		helperFunc()
	}
	helperFunc()

	return fibSequence[len(fibSequence)-1]
}
