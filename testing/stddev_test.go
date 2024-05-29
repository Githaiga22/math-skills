package functions

import (
	"math"
	"testing"
	"math-skills/functions"
)

func TestStandardDeviation(t *testing.T) {
	numbers := []float64{1, 2, 3, 4, 5}
	expected := 1.414213562
	result := functions.StandardDeviation(numbers)
	if math.Abs(result-expected) > 1e-9 {
		t.Errorf("Expected %.9f, got %.9f", expected, result)
	}
}
