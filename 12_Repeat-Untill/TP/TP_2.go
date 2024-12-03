package main

import "fmt"

func main() {

	var angka int

	fmt.Print("Tebak angka (1-10): ")
	fmt.Scan(&angka)

	// Perulangan yang digunakan untuk melakukan tebak angka, dan angka rahasianya adalah 8
	for angka != 8 {
		fmt.Println("Tebakan Anda salah, coba lagi.")
		fmt.Print("Tebak angka (1-10): ")
		fmt.Scan(&angka)
	}

	fmt.Print("Selamat, tebakan Anda benar!")

}