package main

import "fmt"

func main() {

	var angka int

	fmt.Scan(&angka)

	for i := 1; i <= angka; i++ {
		fmt.Println()

		for j := 1; j <= angka; j++ {
			if i == 1 || i == angka || j == 1 || j == angka {
				fmt.Print(i, " ")
			} else {
				fmt.Print("  ")
			}
		}
	}

}
