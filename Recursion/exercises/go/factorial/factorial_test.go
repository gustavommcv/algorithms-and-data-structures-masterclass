package factorial

import "testing"

type testCase struct {
	input int
	want  int
}

func TestFactorial(t *testing.T) {
	tests := []testCase{
		{
			input: 0,
			want:  1,
		},
		{
			input: 2,
			want:  2,
		},
		{
			input: 7,
			want: 5040,
		},
	}

	for _, test := range tests {
		got := factorial(test.input)

		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
