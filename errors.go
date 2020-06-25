// Make a new error type ErrNegativeSqrt that will be returned when the Sqrt
// function is fed a negative number

package main

import (
	"fmt"
	"math"
)

// Minimum change constant
const e = 1e-6

type ErrNegativeSqrt float64

// Make ErrNegativeSqrt an error
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", e)
}

func Sqrt(x float64) (float64, error) {
	// If x is negative, return our custom ErrNegativeSqrt error
	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
	
	var temp float64
	z := float64(1)
	
	for i := 0; i < 10; i++ {
		z, temp = z - (z*z - x) / (2*z), z
		
		// If change in z value is minimal (i.e. less than e constant) return
		if math.Abs(temp - z) < e {
			break
		}
	}
	
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}