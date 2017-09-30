package arbor

import "testing"

func TestAverage(t *testing.T) {
	var v float64
	v = 3 / 2
	if v != 1.5 {
		t.Error("Expected 1.5, got ", v)
	}
}
