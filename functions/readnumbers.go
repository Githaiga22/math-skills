package functions

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadNumbers reads numbers from a file and returns a slice of float64.
func ReadNumbers(filePath string) ([]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue // Skip empty lines
		}
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			continue // Skip lines that can't be parsed to a float
		}
		numbers = append(numbers, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return numbers, nil
}
