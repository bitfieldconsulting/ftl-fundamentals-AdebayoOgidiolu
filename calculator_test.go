package calculator_test

import (
	"calculator"
	"testing"
	"math/rand"
	"time"
	
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		message string
	}

	testCases := []testCase{
		{a: 2, b: 2, want: 4, message: "Two +ve numbers"},
		{a: -3, b: 3, want: 0, message: "One +ve and one -ve number"},
		{a: 5, b: 2, want: 7, message: "Two +ve numbers"},
	}

	for _, tc := range testCases {

		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s : Add(%f , %f): want %f, got %f",tc.message, tc.a, tc.b, tc.want, got)
		}
	}

}

func TestAddRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a + b
		got := calculator.Add(a, b)
		if want != got {
			t.Fatalf("Add(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		message string
	}

	testCases := []testCase {
		{a: 6, b: 5.2, want: 0.7999999999999998, message: "A fractional number in the mix"}, // 0.8 failed the test here. 
		{a: -2, b: 4, want: -6, message: "One -ve and one +ve number"},
		{a: 5, b: 5, want: 0, message: "Two +ve numbers"},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b);
		if tc.want != got {
			t.Errorf("%s : %f minus %f failed. Expected %f but got %f", tc.message, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtractRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a - b
		got := calculator.Subtract(a, b)
		if want != got {
			t.Fatalf("Subtract(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		message string
	}

	testCases := []testCase {
		{a: 7, b: 6, want: 42, message: "One -ve and one +ve number"},
		{a: 6.5, b: 2, want: 13, message: "A fractional number in the mix"},
		{a: 0, b: 3, want: 0, message: "A Zero in the mix"},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s : %f multiplied by %f failed. Expected %f but got %f",tc.message, tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiplyRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := rand.Float64()
		b := rand.Float64()
		want := a * b
		got := calculator.Multiply(a, b)
		if want != got {
			t.Fatalf("Multiply(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}


func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
		errExpected bool
		message string
	}

	testCases := []testCase {
		{a: 7, b: 2, want: 3.5, errExpected: false, message: "Two +ve numbers"},
		{a: 6, b: 0, want: 999, errExpected: true, message: "A divide by zero case"},
		{a: 3, b: 3, want: 1, errExpected: false, message: "Two +ve numbers"},
	}

	for _, tc := range testCases {

		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil

		if errReceived != tc.errExpected {
			t.Fatalf("Divide (%f, %f): unexpected error status: %v", tc.a, tc.b, tc.errExpected)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("%s : %f divided by %f failed. Expected %f but got %f",tc.message, tc.a, tc.b, tc.want, got)
		}

	}
}

func TestDivideRandom(t *testing.T) {
	t.Parallel()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		a := rand.Float64()
		b := rand.Float64()
		if b == 0 {
			t.Fatal("Divide by zero not allowed!")
		}
		want := a/b
		got, err := calculator.Divide(a, b)
		if err != nil{
			t.Error("Division failed.")
		}
		if want != got {
			t.Fatalf("Divide(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}


func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		name        string
		input       float64
		want        float64
		errExpected bool
	}
	testCases := []testCase{
		{name: "Square root of 4 is 2", input: 4, want: 2, errExpected: false},
		{name: "Square root of one is one", input: 1, want: 1, errExpected: false},
		{name: "Square root of a negative number is not allowed (unless you're a physicist)", input: -1, want: 999, errExpected: true},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.input)
		errReceived := err != nil
		if errReceived != tc.errExpected {
			t.Fatalf("%s: Sqrt(%f): unexpected error status: %v", tc.name, tc.input, err)
		}
		if !tc.errExpected && tc.want != got {
			t.Errorf("%s: Sqrt(%f): want %f, got %f", tc.name, tc.input, tc.want, got)
		}
	}
}