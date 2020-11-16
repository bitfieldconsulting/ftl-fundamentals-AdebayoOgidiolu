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

// Attempting to break the test
// a return type of float32 returns "cannot use a * b (type float64) as type float32 in return argument"
// changing the data type for a to bool returns "invalid operation: a * b (mismatched types bool and float64)""
func Multiply(a, b float64) float64 {  
	return a * b
}


func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("bad inputs %f, %f: division by zero is not allowed", a, b)
	}
	return a / b, nil
}


func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("bad input %f: square root of a negative number is not allowed", a)
	}
	return math.Sqrt(a), nil
}

func AddMany(inputs ...float64) float64{
    var sum float64
    for _,input := range inputs{
        sum = sum + input
        
    }
    return sum
}

func SubtractMany(inputs ...float64) float64{
    var total float64
    for _,input := range inputs{
        total = total - input
        
    }
    return total
}

func MultiplyMany(inputs ...float64) float64{
    var total float64
    for index,input := range inputs{
        if index == 0{
            total = inputs[index]
        }else{
            total = total * input
        }
        
    }
	return total
	
}

func DivideMany(inputs ...float64) (float64){
    var val float64
    
    for index,input := range inputs{
        if input == 0{
            fmt.Println("Divide by Zero not allowed")
        }
        if index == 0{
            val = inputs[index]
        }else{
            val = val/input
        }
        
    }
    return val
}