package main

import "fmt"

func main() {

	var bunga, pita string
	var jumlah int

	jumlah = 0

	for {

		fmt.Print("Bunga ", jumlah+1, ": ")
		fmt.Scan(&bunga)

		// ketika mengetikan selesai, maka program selesai
		if bunga == "selesai" {
			break
		}

		// Digunakan untuk menambahkan setiap masukan yang dimasukan
		pita = pita + bunga + " - "
		jumlah++

	}

	fmt.Println("Pita:", pita)
	fmt.Println("Bunga:", jumlah)

}