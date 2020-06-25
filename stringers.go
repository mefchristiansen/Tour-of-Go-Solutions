// Make the IPAddr type implement fmt.Stringer to print the address as a
// dotted quad

package main

import "fmt"

type IPAddr [4]byte

// Have the IPAddr type implement fmt.Stringer
func (addr IPAddr) String() string {
	var quad string
	
	quad += fmt.Sprintf("%v", addr[0])
	
	for i:= range addr[1:] {
		quad +=	fmt.Sprintf(".%v", addr[i + 1])
	}
	
	return quad
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}