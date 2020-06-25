// Implement Pic, a 2D uint8 array of size dy * dx

// Each element of a slice of length dy is a slice of dx 8-bit
// unsigned integers

// This will display a picture

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Initialize a 2D uint8 array of size dy
	pic := make([][]uint8, dy)

	for x := range pic {
		// Intialize each row in pic of size dx
		pic[x] = make([]uint8, dx)

		for y := range pic[x] {
			pic[x][y] = uint8(x ^ y)
		}
	}

	return pic
}

func main() {
	pic.Show(Pic)
}