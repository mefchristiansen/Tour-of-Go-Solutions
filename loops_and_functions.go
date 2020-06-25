// Given a number x, find the number z for which z² is most nearly x.

// Computers typically compute the square root of x using the following loop:
// z -= (z*z - x) / (2*z)

// Repeating this adjustment makes the guess better and better until we reach
// an answer that is as close to the actual square root as can be

package main

import (
	"fmt"
	"math"
)

// Minimum change constant
const e = 1e-6

func Sqrt(x float64) float64 {
	var temp float64
	z := float64(1)
	
	for i := 0; i < 10; i++ {
		z, temp = z - (z*z - x) / (2*z), z
		
		// If change in z value is minimal (i.e. less than e constant) return
		if math.Abs(temp - z) < e {
			break
		}
		
		fmt.Println(z)
	}
	
	return z
}

func main() {
	estimate := Sqrt(2)
	fmt.Printf("Estimate: %v, Library: %v, Difference: %v",
		estimate, math.Sqrt(2),
		math.Abs(estimate - math.Sqrt(2)),
	)
}