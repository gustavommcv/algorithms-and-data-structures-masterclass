package power

import "testing"

func TestPower(t *testing.T) {
	base := 4
	exponent := 3
	expected := 64

	got := Power(base, exponent)

	if got != expected {
		t.Errorf("got %v, expected %v", got, expected)
	}
}
