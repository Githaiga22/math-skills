package functions

import (
	"os"
	"testing"
	"math-skills/functions"
)

// TestReadNumbers tests the ReadNumbers function.
func TestReadNumbers(t *testing.T) {
	// Content to be written to the temporary file for testing
	content := "1.0\n2.0\n3.0\n4.0\n5.0\n"

	// Create a temporary file
	tmpfile, err := os.CreateTemp("", "testdata")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(tmpfile.Name()) // Ensure the file is cleaned up

	// Write the content to the temporary file
	if _, err := tmpfile.WriteString(content); err != nil {
		t.Fatal(err)
	}

	// Close the file to ensure the written content is flushed
	if err := tmpfile.Close(); err != nil {
		t.Fatal(err)
	}

	// Define the expected result
	expected := []float64{1.0, 2.0, 3.0, 4.0, 5.0}

	// Call the ReadNumbers function with the temporary file path
	result, err := functions.ReadNumbers(tmpfile.Name())
	if err != nil {
		t.Fatalf("Error reading numbers: %v", err)
	}

	// Check if the result length matches the expected length
	if len(result) != len(expected) {
		t.Fatalf("Expected %d numbers, got %d", len(expected), len(result))
	}

	// Check if each number in the result matches the expected number
	for i, num := range expected {
		if result[i] != num {
			t.Errorf("Expected %.2f, but got %.2f", num, result[i])
		}
	}
}
