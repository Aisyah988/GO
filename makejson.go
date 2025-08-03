package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Prompt for name
	fmt.Print("Enter your name: ")
	nameInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read name input:", err)
		return
	}
	name := strings.TrimSpace(nameInput)

	// Prompt for address
	fmt.Print("Enter your address: ")
	addressInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read address input:", err)
		return
	}
	address := strings.TrimSpace(addressInput)

	// Create map
	person := map[string]string{
		"name":    name,
		"address": address,
	}

	// Marshal to JSON
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Println("Failed to marshal JSON:", err)
		return
	}

	// Print JSON
	fmt.Println(string(jsonData))
}
