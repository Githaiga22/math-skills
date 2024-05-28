package functions

import "math"

// StandardDeviation calculates the standard deviation of a list of numbers.
func StandardDeviation(numbers []float64) float64 {
	return math.Sqrt(Variance(numbers))
}
