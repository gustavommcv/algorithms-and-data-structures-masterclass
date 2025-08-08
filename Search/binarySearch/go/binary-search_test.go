package binary_search

import "testing"

func TestIndexOf(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		input := 4
		want := 3

		got := binarySearch(arr, input)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("success", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		input := 2
		want := 1

		got := binarySearch(arr, input)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("not found", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		input := 7
		want := -1

		got := binarySearch(arr, input)

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
