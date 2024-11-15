package main

import (
	"fmt"
)

func main() {
	var berat int

	fmt.Print("Masukan Berat: ")
	fmt.Scanln(&berat)

	beratBarang := berat / 1000
	gram := berat % 1000
	harga := beratBarang * 10000

	if beratBarang > 10 {
		fmt.Println("Detail Berat adalah: ", beratBarang, "Kg")
		fmt.Println("Detail Berat adalah: ", gram, "gram")
		fmt.Println("Harga total adalah Rp. ", harga)

	} else if beratBarang <= 10 && gram < 500 {
		sisaGram := gram * 15
		fmt.Println("Detail Berat adalah: ", beratBarang, "Kg")
		fmt.Println("Detail Berat adalah: ", gram, "gram")
		fmt.Println("Harga totalnya adalah: ", sisaGram+harga)
	} else if beratBarang <= 10 && gram >= 500 {
		sisaGram := gram * 5
		fmt.Println("Detail Berat adalah: ", beratBarang, "Kg")
		fmt.Println("Detail Berat adalah: ", gram, "gram")
		fmt.Println("Harga totalnya adalah: ", sisaGram+harga)
	} else {
		fmt.Println("GBLK")
	}
}