package main

import (
	"fmt"
)

func main() {
	var n int

	// Meminta input dari pengguna
	fmt.Print("Masukkan bilangan bulat ganjil: ")
	fmt.Scan(&n)

	// Memeriksa apakah n adalah bilangan ganjil
	if n%2 == 0 {
		fmt.Println("Input harus bilangan bulat ganjil.")
		return
	}

	// Mencetak pola X dengan angka
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if j == i || j == n-i-1 {
				fmt.Print(i + 1) // Mencetak angka sesuai dengan baris (i + 1)
			} else {
				fmt.Print(" ") // Mencetak spasi
			}
		}
		fmt.Println() // Pindah ke baris berikutnya
	}
}