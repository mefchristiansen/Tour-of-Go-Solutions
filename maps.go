// Return a map of of the counts of each word in the string s

package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	wordCount := make(map[string]int)
	
	// Iterate through each word in the string s
	for _, v := range strings.Fields(s) {
		// Increment the count of each word in the map
		wordCount[v]++
	}
		
	return wordCount
}

func main() {
	wc.Test(WordCount)
}