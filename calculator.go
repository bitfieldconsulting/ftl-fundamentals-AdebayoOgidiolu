// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"fmt"
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

func AddMany(a, b float64, inputs ...float64) float64 {
	sum := a + b
	for _, input := range inputs {
		sum += input

	}
	return sum
}

func SubtractMany(a, b float64, inputs ...float64) float64 {
	sum := a - b
	for _, input := range inputs {
		sum -= input

	}
	return sum
}

func MultiplyMany(a, b float64, inputs ...float64) float64 {
	sum := a * b
	for _, input := range inputs {
		sum *= input
	}
	return sum

}

func DivideMany(a, b float64, inputs ...float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("bad input: Divide by zero is not allowed")
	}
	sum := a / b

	for _, input := range inputs {
		if input == 0 {
			return 0, fmt.Errorf("bad input: Divide by zero is not allowed")
		}
		sum /= input

	}
	return sum, nil
}
