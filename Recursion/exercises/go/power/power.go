package power

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

func Power[T Number](base, exponent T) T {
	if exponent == 0 {
		return 1
	}

	return base * Power(base, exponent-1)
}

// func Power[T Number](base, exponent T) T {
// 	var result T = 1
//
// 	for i := 0; i < int(exponent); i++ {
// 		result *= base
// 	}
//
// 	return result
// }
