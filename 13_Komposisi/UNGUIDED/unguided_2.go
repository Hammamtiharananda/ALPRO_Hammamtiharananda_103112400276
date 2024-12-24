package main

import "fmt"

func main() {
	var angka int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// Ketika nilai angka kurang dari atau sama dengan 1, maka program langsung berhenti
	if angka <= 1 {
		fmt.Println("Bukan prima")
		return
	}

	// Perulangan untuk memeriksa apakah angka tersebut prima
	for i := 2; i < angka; i++ {
		if angka%i == 0 {
			fmt.Println("Bukan prima")
			return
		}
	}

	fmt.Println("Prima")
}