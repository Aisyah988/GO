package main

import "fmt"

func main() {
	var input float32

	// Prompt the user to enter a floating-point number
	fmt.Print("Enter a ploting point numbers :")
	_, err := fmt.Scanf("%f", &input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid floating point number.")
		return
	}
	// Convert the float to an integer by truncation
	truncated := int(input)

	// Print the truncated result
	fmt.Println("Truncated (integer) value:", truncated)
}
