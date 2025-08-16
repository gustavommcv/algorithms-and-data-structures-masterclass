package product_of_array

import "testing"

type testCase struct {
	input []int
	want  int
}

func TestFactorial(t *testing.T) {
	tests := []testCase{
		{
			input: []int{1, 2, 3},
			want:  6,
		},
		{
			input: []int{1, 2, 3, 10},
			want:  60,
		},
	}

	for _, test := range tests {
		got := ProductOfArray(test.input)

		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
