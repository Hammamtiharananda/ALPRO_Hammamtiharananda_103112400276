package main

import (
	"fmt"
)

func main() {

	// Variabel untuk menyimpan data
	var (
		n int
	)

	// Inputan untuk memasukan nilai n
	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

	// Perulangan untuk menampilkan angka 
	for i := 1; i < n; i++ {
		// Mengalikan angka perulangan dengan angka itu sendiri (1*1 = 1, 2*2 = 4 ... nilai n)
		fmt.Printf("%d ", i*i)
	}

	// Digunakan untuk menampilkan nilai n terakhir
	fmt.Print(n * n)
}