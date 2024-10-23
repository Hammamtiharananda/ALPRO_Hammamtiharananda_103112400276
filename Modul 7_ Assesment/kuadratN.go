package main

import (
	"fmt"
)
func main() {
	var N int

	// Input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	// Validasi input
	if N <= 0 {
		fmt.Println("Error: Masukkan harus bilangan bulat positif.")
		return
	}

	// Mencetak kuadrat dari 1 sampai N
	for i := 1; i <= N; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println() // Menambahkan newline setelah output
}