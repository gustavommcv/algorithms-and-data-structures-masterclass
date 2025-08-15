package main

import (
	"reflect"
	"testing"
)

func TestCollectOddValues(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expect := []int{1, 3, 5}

	got := collectOddValues(input)

	if !reflect.DeepEqual(got, expect) {
		t.Errorf("got %v, expect %v", got, expect)
	}
}
