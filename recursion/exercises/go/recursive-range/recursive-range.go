package recursive_range

func RecursiveRange(n int) int {
	if n < 1 {
		return 0
	}

	return n + RecursiveRange(n-1)
}
