// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
	"log"
	"math"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a, b float64) float64 {
	return a + b
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a, b float64) float64 {
	return a - b

}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input: Divide by zero is not allowed")
	}
	return a / b, nil
}

func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input %f: square root of a negative number is not allowed", a)
	}
	return math.Sqrt(a), nil
}

func AddMany(inputs ...float64) float64 {
	var sum float64
	if len(inputs) < 2 {
		log.Fatalf("Please specify at least 2 values.")
	}
	for _, input := range inputs {
		sum += input

	}
	return sum
}

func SubtractMany(inputs ...float64) float64 {
	var total float64
	for _, input := range inputs {
		total = total - input

	}
	return total
}

func MultiplyMany(inputs ...float64) float64 {
	var total float64
	for index, input := range inputs {
		if index == 0 {
			total = inputs[index]
		} else {
			total = total * input
		}

	}
	return total

}

func DivideMany(inputs ...float64) (float64, error) {
	var val float64

	for index, input := range inputs {
		if input == 0 {
			return 0, fmt.Errorf("bad input: Divide by zero is not allowed")
		}
		if index == 0 {
			val = inputs[index]
		} else {
			val = val / input
		}

	}
	return val, nil
}

func CloseEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}
