package bubbleSort

import (
	"reflect"
	"testing"
)

type testCase struct {
	input []int
	want  []int
}

func TestFactorial(t *testing.T) {
	tests := []testCase{
		{
			input: []int{5, 4, 3, 2, 1},
			want:  []int{1, 2, 3, 4, 5},
		},
	}

	for _, test := range tests {
		BubbleSort(test.input)

		if !reflect.DeepEqual(test.want, test.input) {
			t.Errorf("got %v, want %v", test.input, test.want)
		}
	}
}
