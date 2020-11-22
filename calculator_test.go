package calculator_test

import (
	"calculator"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name string
		a, b float64
		want float64
	}{
		{
			name: "Two plus two equals four",
			a:    2,
			b:    2,
			want: 4,
		},
		{
			name: "Negative three plus three equals zero",
			a:    -3,
			b:    3,
			want: 0,
		},
		{
			name: "Five plus two equals Seven",
			a:    5,
			b:    2,
			want: 7,
		},
	}

	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s : Add(%f , %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
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
	testCases := []struct {
		name string
		a, b float64
		want float64
	}{
		{
			name: "Six minus five point two equals zero point eight",
			a:    6,
			b:    5.2,
			want: 0.8,
		},
		{
			name: "Negative two minus four equals -negative six",
			a:    -2,
			b:    4,
			want: -6,
		},
		{
			name: "Five minus five equals zero",
			a:    5,
			b:    5,
			want: 0,
		},
	}

	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if !CloseEnough(tc.want, got) {
			t.Errorf("%s : Add(%f , %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
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
	testCases := []struct {
		name    string
		a, b    float64
		want    float64
		message string
	}{
		{
			name: "Seven multiplied by six is forty two",
			a:    7,
			b:    6,
			want: 42,
		},
		{
			name: "Six point five multiplied by two is thirteen",
			a:    6.5,
			b:    2,
			want: 13,
		},
		{
			name: "Zero multiplied by three is zero",
			a:    0,
			b:    3,
			want: 0,
		},
	}

	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("%s : Multiply(%f , %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
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

func CloseEnough(a, b float64) bool {
	tolerance := 0.0000002
	return math.Abs(a-b) <= tolerance
}

func TestDivide(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		a, b        float64
		want        float64
		errExpected bool
	}{
		{
			name:        "Seven divided by two is three point five",
			a:           7,
			b:           2,
			want:        3.5,
			errExpected: false,
		},
		{
			name:        "Six divided by zero is nine hundred and ninety nine",
			a:           6,
			b:           0,
			want:        999,
			errExpected: true,
		},
		{
			name:        "Three divided by three is one",
			a:           3,
			b:           3,
			want:        1,
			errExpected: false,
		},
	}

	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)
		errReceived := err != nil

		if errReceived != tc.errExpected {
			t.Fatalf("%s :Divide (%f, %f): unexpected error status: %v", tc.name, tc.a, tc.b, tc.errExpected)
		}

		if !tc.errExpected && tc.want != got {
			t.Errorf("%s : Divide(%f , %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
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
			continue
		}
		want := a / b
		got, err := calculator.Divide(a, b)
		if err != nil {
			t.Error("Division failed.")
		}
		if want != got {
			t.Fatalf("Divide(%f, %f): want %f, got %f", a, b, want, got)
		}
	}
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		input       float64
		want        float64
		errExpected bool
	}{

		{
			name:        "Square root of 4 is 2",
			input:       4,
			want:        2,
			errExpected: false,
		},
		{
			name:        "Square root of one is one",
			input:       1,
			want:        1,
			errExpected: false,
		},
		{
			name:        "Square root of a negative number is not allowed (unless you're a physicist)",
			input:       -1,
			want:        999,
			errExpected: true,
		},
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

func TestAddMany(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name  string
		a, b  float64
		extra []float64
		want  float64
	}{
		{
			name: "Two plus two equals four",
			a:    2,
			b:    2,
			want: 4},
		{
			name: "Negative three plus three equals zero",
			a:    -3,
			b:    3,
			want: 0,
		},
		{
			name: "Five plus two equals Seven",
			a:    5,
			b:    2,
			want: 7,
		},
		{
			name:  "Supply extra inputs",
			a:     5,
			b:     2,
			extra: []float64{2, 3},
			want:  12,
		},
	}

	for _, tc := range testCases {
		got := calculator.AddMany(tc.a, tc.b, tc.extra...)
		if tc.want != got {
			t.Errorf("%s : AddMany(%f , %f): want %f, got %f", tc.name, tc.a, tc.b, tc.want, got)
		}
	}

}
