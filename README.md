# Math Skills Project

## Overview

The Math Skills project is designed to calculate four key statistical measures: Average, Median, Variance, and Standard Deviation from a dataset provided in a file. This project is written in Go and demonstrates fundamental concepts in statistics and programming.

## Objectives

- Calculate the following statistics from a dataset:
  - Average
  - Median
  - Variance
  - Standard Deviation

## File Structure

The project directory is organized as follows:

A directory named math-skills that contains ; - main.go | README.md | go.mod | data.txt and a sub-directory - functions.Within the functions sub-directory contains ; - average.go | average_test.go | median.go | median_test.go | readnumbers.go | readnumbers_test.go | stddev.go | stddev_test.go | variance.go | variance_test.go .


### Description of Files

- **functions/**: Contains individual Go files for each statistical calculation and their corresponding tests.
  - `average.go`: Calculates the average.
  - `average_test.go`: Tests for `average.go`.
  - `median.go`: Calculates the median.
  - `median_test.go`: Tests for `median.go`.
  - `readnumbers.go`: Reads numbers from the file.
  - `readnumbers_test.go`: Tests for `readnumbers.go`.
  - `stddev.go`: Calculates the standard deviation.
  - `stddev_test.go`: Tests for `stddev.go`.
  - `variance.go`: Calculates the variance.
  - `variance_test.go`: Tests for `variance.go`.
- `main.go`: The main entry point of the program.
- `go.mod`: Go module file.
- `README.md`: Project documentation.
- `data.txt`: The data file containing numbers, one per line.

## Instructions

### Prerequisites

Ensure you have Go installed on your machine. You can download it from the [official Go website](https://golang.org/dl/).

### Running the Program

1. **Place your data file in the project root directory**. The file should be named `data.txt` and contain one number per line. An example `data.txt` file with random numbers is provided.

2. **Navigate to the project root directory in your terminal**.

    ```sh
    cd path/to/your/math-skills
    ```

3. **Run the program**:

    ```sh
    go run main.go data.txt
    ```

    This command will read the numbers from `data.txt` and print the calculated statistics to the console.

### Example Output

Average: 1000000

Median: 1000000

Variance: 333333333333

Standard Deviation: 57735

### Testing

To run tests for the functions:

1. **Navigate to the project root directory** in your terminal.

2. **Run the tests**:

    ```sh
    go test ./functions
    ```

    This command will run all tests in the `functions` package and display the results.

## Error Handling

- If the file path is not provided as an argument, the program will log an error and exit.
- If there's an issue reading the data from the file (e.g., file not found, incorrect format), the program will log an error and exit.
   - Fewer Than 2 Numbers: Ensures the program exits with an appropriate error message if the data file contains fewer than 2 numbers.
   -  Empty Lines/Whitespaces: The ReadNumbers function trims and skips empty lines.
   - Decimal Values: The ReadNumbers function now parses floating-point numbers.
    
    

- Ensure that `data.txt` is correctly formatted with one number per line.
- The program rounds the calculated statistics to the nearest integer before printing.
- The program and test files follow good coding practices for readability and maintainability.

## Possible Errors

- **File Not Found**: Ensure the `data.txt` file is in the correct directory.
- **Incorrect File Format**: Ensure `data.txt` contains only numbers, one per line.
- **Missing Arguments**: Provide the file path as an argument when running the program.

## Conclusion

The Math Skills project is a practical application to understand basic statistical calculations and Go programming. By following the instructions and using the provided structure, you can calculate key statistics from any dataset provided in a file.
