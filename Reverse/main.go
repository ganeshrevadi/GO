package main

import "fmt"

func reverseString(s string) string {
	// Convert string to slice of runes
	runes := []rune(s)

	// Iterate over slice and swap first and last elements
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert slice of runes back to string
	return string(runes)
}

func main() {
	s := "Hello, world!"
	fmt.Println(reverseString(s))
}
