package factorial

type Number interface {
	~int | ~float64
}

func factorial[T Number](n T) T {
	if n <= 1 {
		return 1
	}

	return n * factorial(n - 1)
}
