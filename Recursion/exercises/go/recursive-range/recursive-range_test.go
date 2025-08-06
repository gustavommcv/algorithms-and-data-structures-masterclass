package recursive_range

import "testing"

type testCase struct {
	input int
	want  int
}

func TestFactorial(t *testing.T) {
	tests := []testCase{
		{
			input: 6,
			want:  21,
		},
		{
			input: 10,
			want:  55,
		},
	}

	for _, test := range tests {
		got := RecursiveRange(test.input)

		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
