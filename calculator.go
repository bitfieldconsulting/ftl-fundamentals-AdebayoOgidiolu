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

// Multiply takes two numbers and returns the result of multiplying the first
// and the second
func Multiply(a, b float64) float64 {
	return a * b
}

// Divide takes two numbers and returns the result of dividing the first
// by the second. An error is returned if operation divide fails
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad input: Divide by zero is not allowed")
	}
	return a / b, nil
}

// Sqrt takes a non-negative number and returns the square root
// It returns an error if the number is less than zero
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input %f: square root of a negative number is not allowed", a)
	}
	return math.Sqrt(a), nil
}

// AddMany takes at least two numbers and adds them all together
func AddMany(a, b float64, inputs ...float64) float64 {
	sum := a + b
	for _, input := range inputs {
		sum += input
	}
	return sum
}

// SubtractMany takes at least two numbers and subtracts one number
// from the previous until all arguments have been used
func SubtractMany(a, b float64, inputs ...float64) float64 {
	sum := a - b
	for _, input := range inputs {
		sum -= input
	}
	return sum
}

// MultiplyMany takes at least two numbers and multiplies one number
// with the next until all arguments have been used
func MultiplyMany(a, b float64, inputs ...float64) float64 {
	sum := a * b
	for _, input := range inputs {
		sum *= input
	}
	return sum
}

// DivideMany takes at least two numbers and divides one number
// by the next until all arguments have been used
// Arg 'b' cannot be a zero
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
