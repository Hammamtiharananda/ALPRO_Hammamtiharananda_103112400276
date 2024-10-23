package main

import (
	"fmt"
)

func jumlahBilangan(x, y int) int {
	total := 0
	for i := x; i <= y; i++ {
		total += i
	}
	return total
}
func main() {
	var x, y int

	// input dari pengguna
	fmt.Print("Masukkan bilangan bulat positif x (x <= y): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan bulat positif y (x <= y): ")
	fmt.Scan(&y)

	// Validasi input
	if x <= y {
		hasil := jumlahBilangan(x, y)
		fmt.Printf("Hasil penjumlahan dari %d sampai %d adalah: %d\n", x, y, hasil)
	}else 
	{
		fmt.Println("Pastikan x selalu lebih kecil atau sama dengan y.")
	}
}
