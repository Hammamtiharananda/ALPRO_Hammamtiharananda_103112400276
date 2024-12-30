package main

import "fmt"

func main() {

	var angka int

	fmt.Print("Masukkan angka : ")
	fmt.Scan(&angka)

	if angka%3 == 0 && angka%5 == 0 {
		fmt.Println("Kelipatan 3 dan 5")
	} else if angka%3 == 0 {
		fmt.Println("Kelipatan 3")
	} else if angka%5 == 0 {
		fmt.Println("Kelipatan 5")
	} else {
		fmt.Println("Bukan kelipatan 3 dan 5")
	}

}
