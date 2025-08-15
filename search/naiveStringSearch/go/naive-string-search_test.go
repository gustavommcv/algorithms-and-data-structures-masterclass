package naiveStringSearch

import "testing"

type testCase struct {
	str    string
	target string
	want   int
}

func TestFactorial(t *testing.T) {
	tests := []testCase{
		{
			str:    "wowomgzomg",
			target: "omg",
			want:   2,
		},
		{
			str:    "wowomgzomg",
			target: "ferr",
			want:   0,
		},
		{
			str:    "lorie loled",
			target: "lol",
			want:   1,
		},
	}

	for _, test := range tests {
		got := naiveStringSearch(test.str, test.target)

		if got != test.want {
			t.Errorf("got %v, want %v", got, test.want)
		}
	}
}
