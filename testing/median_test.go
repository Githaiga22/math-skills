package functions

import (
	"testing"
	"math-skills/functions"
)

func TestMedian(t *testing.T) {
	numbers := []float64{1, 3, 3, 6, 7, 8, 9}
	expected := 6.0
	result := functions.Median(numbers)
	if result != expected {
		t.Errorf("Expected %.1f, got %.1f", expected, result)
	}
}
