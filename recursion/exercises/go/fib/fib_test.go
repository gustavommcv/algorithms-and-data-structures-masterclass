package fib

import "testing"

type testCase struct {
	input int
	want  int
}

func TestFib(t *testing.T) {
	tests := []testCase{
		{
			input: 4,
			want:  3,
		},
		{
			input: 10,
			want:  55,
		},
		{
			input: 28,
			want:  317811,
		},
		{
			input: 35,
			want:  9227465,
		},
	}

	for _, test := range tests {
		got := fib(test.input)

		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}

}
