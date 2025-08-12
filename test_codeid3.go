package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan angka: ")
	fmt.Scan(&n)

	// Buat pola 1 baris
	var pola string
	for i := 2; i <= n; i += 2 {
		pola += fmt.Sprintf("-%d", i)
	}
	for i := 1; i <= n; i += 2 {
		pola += fmt.Sprintf("%d", i)
	}

	// Cetak n baris
	for i := 0; i < n; i++ {
		fmt.Println(pola)
	}
}
