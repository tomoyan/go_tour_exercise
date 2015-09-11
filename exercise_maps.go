/*
Exercise: Maps
Implement WordCount.
It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/
package main

import (
	"fmt"
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	// fmt.Println("string is:", s)
	// fmt.Printf("Fields are: %q\n", strings.Fields(s))

	m := make(map[string]int)
	a := strings.Fields(s)

	for i := range a {
		m[a[i]] += 1
	}

	// fmt.Println("array:", a)
	// fmt.Println("map:", m)
	// return map[string]int{"x": 1}
	return m
}

func main() {
	fmt.Println("### Start Map Exercise ###")

	wc.Test(WordCount)

	fmt.Println("### End Map Exercise ###")
}
