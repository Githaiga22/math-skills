package functions

import "testing"

func TestVariance(t *testing.T) {
	numbers := []float64{1, 2, 3, 4, 5}
	expected := 2.0
	result := Variance(numbers)
	if result != expected {
		t.Errorf("Expected %.1f, got %.1f", expected, result)
	}
}
