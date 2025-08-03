package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Create an empty slice of integers with initial capacity 3
	numbers := make([]int, 0, 3)

	reader := bufio.NewReader(os.Stdin)

	for {
		// Prompt user for input
		fmt.Print("Enter an integer or 'X' to exit: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Check if user wants to exit
		if strings.ToUpper(input) == "X" {
			fmt.Println("Exiting program.")
			break
		}

		// Try to convert input to integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'X' to exit.")
			continue
		}

		// Append number to the slice
		numbers = append(numbers, num)

		// Sort the slice
		sort.Ints(numbers)

		// Print the sorted slice
		fmt.Println("Sorted slice:", numbers)
	}
}
