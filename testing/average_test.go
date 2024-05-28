package functions

import "testing"

func TestAverage(t *testing.T) {
	numbers := []float64{1, 2, 3, 4, 5}
	expected := 3.0
	result := Average(numbers)
	if result != expected {
		t.Errorf("Expected %.1f, got %.1f", expected, result)
	}
}
