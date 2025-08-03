package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Struct with first and last name
type Name struct {
	fname string
	lname string
}

func main() {
	// Prompt for file name
	fmt.Print("Enter the file name: ")
	var filename string
	fmt.Scanln(&filename)

	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var names []Name
	scanner := bufio.NewScanner(file)

	// Read file line by line
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			fname := truncate(parts[0], 20)
			lname := truncate(parts[1], 20)
			names = append(names, Name{fname, lname})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print all names
	for _, n := range names {
		fmt.Printf("First Name: %s, Last Name: %s\n", n.fname, n.lname)
	}
}

// Helper function to limit string length to 20 characters
func truncate(s string, maxLen int) string {
	if len(s) > maxLen {
		return s[:maxLen]
	}
	return s
}
