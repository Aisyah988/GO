package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string

	// Prompt the user to enter a string character
	fmt.Print("insert a character string containing the letters i, a, and n :")
	fmt.Scanln(&input)
	// Convert to lowercase to make the check case-insensitive
	lowerInput := strings.ToLower(input)

	// Check if it starts with 'i', contains 'a', and ends with 'n'
	startsWithI := strings.HasPrefix(lowerInput, "i")
	containsA := strings.Contains(lowerInput, "a")
	endsWithN := strings.HasSuffix(lowerInput, "n")

	// Print result
	if startsWithI && containsA && endsWithN {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
