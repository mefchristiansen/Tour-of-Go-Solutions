// Implement a rot13Reader (https://en.wikipedia.org/wiki/ROT13) that
// implements io.Reader and reads from an io.Reader

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot *rot13Reader) Read(b []byte) (int, error) {
	// Read from io.Reader
	n, err := rot.r.Read(b)
	
	for i := range b {
		if b[i] < 'A' || b[i] > 'z' { //Ignore non-alphabetical characters
            continue
		} else if (b[i] > 'a' && b[i] < 'a' + 13) ||
			(b[i] > 'A' && b[i] < 'A' + 13) {
			// If character at index i is one of first 13 characters in the
			// alphabet, advance it 13 characters

			b[i] += 13
		} else {
			// Else, withdraw it 13 characters

			b[i] -= 13
		}
	}
	
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}