// Calculate the fibonacci sequence using a closure

package main

import "fmt"

// fibonacci is a function that returns
// a function (i.e. a closure) that returns an int.
func fibonacci() func() int {
	a := 0
	b := 1
	
	return func() int {	
		a, b = b, a + b
		
		return b - a
	}
}

func main() {
	f := fibonacci()
	// Calculate first 10 numbers in fibonacci sequence (zero based)
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}