package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			if j%2 != 0 {
				fmt.Print(i)
			} else {
				fmt.Print(n - i + 1)
			}
		}
		fmt.Println()
	}
}
