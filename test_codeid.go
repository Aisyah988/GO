package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		if i%2 == 0 {
			// Baris ganjil (index 0, 2, 4...): dari n ke 1
			for j := n; j >= 1; j-- {
				fmt.Print(j)
			}
		} else {
			// Baris genap (index 1, 3, 5...): dari 1 ke n
			for j := 1; j <= n; j++ {
				fmt.Print(j)
			}
		}
		fmt.Println()
	}
}
