package functions

import (
	"testing"

	"math-skills/functions"
)

func TestVariance(t *testing.T) {
	numbers := []float64{1, 2, 3, 4, 5}
	expected := 2.0
	result := functions.Variance(numbers)
	if result != expected {
		t.Errorf("Expected %.1f, got %.1f", expected, result)
	}
}
