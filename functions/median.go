package functions

import "sort"

// Median calculates the median of a list of numbers.
func Median(numbers []float64) float64 {
	sort.Float64s(numbers)
	n := len(numbers)
	if n%2 == 0 {
		return float64(numbers[n/2-1]+numbers[n/2]) / 2.0
	}
	return float64(numbers[n/2])
}
