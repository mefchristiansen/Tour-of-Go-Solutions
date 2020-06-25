// Determine if two binary trees store the same values

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	// Declare nested recursive function to allow channel to be
	// closed in Walk function
	var walker func (t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		
    	// In order traversal
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	
	walker(t)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	
	// Iterate through received values from channel 1
	for v1 := range ch1 {
		v2, ok := <- ch2
		
		// Return false if there are no more values from channel 2
		// (i.e. difference in number of values), or if the values
		// from the two channels don't match up
		if !ok || v1 != v2 {
			return false
		}
	}
	
	// Need to check if channel 2 has any values remaining.
	// If it does, then there is difference in number of values
	// between the two trees.
	_, ok := <- ch2
	return !ok
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}