package functions

import "math"

// Variance calculates the variance of a list of numbers.
func Variance(numbers []float64) float64 {
	mean := Average(numbers)
	sumOfSquares := 0.0
	for _, num := range numbers {
		sumOfSquares += math.Pow(float64(num)-mean, 2)
	}
	return sumOfSquares / float64(len(numbers))
}
