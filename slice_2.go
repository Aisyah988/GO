package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	instr := ""
	ia := make([]int, 0, 3)

	for {
		fmt.Printf("Input integer: ")
		fmt.Scanln(&instr)

		if instr == "X" {
			break
		}

		num, err := strconv.Atoi(instr)
		if err == nil {
			ia = append(ia, num)
			sort.Ints(ia)
		} else {
			fmt.Println("Invalid input. Enter a number or 'X' to exit.")
		}

		fmt.Println(ia)
	}
}
