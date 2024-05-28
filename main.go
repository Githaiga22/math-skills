package main

import (
	"fmt"
	"log"
	"math"
	"os"

	"math-skills/functions" // Correct import path
)

func main() {
	// Ensure the file path is provided as an argument
	if len(os.Args) != 2 {
		log.Fatalf("provide extra argument")
	}

	// Read data from the file
	filePath := os.Args[1]
	data, err := functions.ReadNumbers(filePath)
	if err != nil {
		log.Fatalf("Error reading data: %v", err)
	}

	// Handle the case where data has fewer than 2 numbers
	if len(data) < 2 {
		log.Fatalf("Data file must contain at least two numbers")
	}

	// Calculate and print statistics
	average := functions.Average(data)
	median := functions.Median(data)
	variance := functions.Variance(data)
	stdDev := functions.StandardDeviation(data)

	// Print the results rounded to the nearest whole number
	fmt.Printf("Average: %d\n", int(math.Round(average)))
	fmt.Printf("Median: %d\n", int(math.Round(median)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))

	os.Exit(0) // Indicating successful completion
}
