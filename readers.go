// Implement a Reader type that emits an infinite stream of the ASCII
// character 'A'.

package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(b []byte) (int, error) {
	// Populate given byte slice b with ASCII character 'A'
	for i := range b {
		b[i] = 'A'
	}
	
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}