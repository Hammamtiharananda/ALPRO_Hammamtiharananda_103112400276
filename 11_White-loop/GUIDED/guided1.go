package main

import "fmt"

func main() {

	// Membuat variabel
	var n, j int

	// Inputan untuk nilai n
	fmt.Print("Masukan Nilai n: ")
	fmt.Scan(&n)

	j = n

	// Perulangan menggunakan while loop
	for j > 1 {
		fmt.Print(j, " x ")
		j = j - 1
	}

	// Menampilkan hasil
	fmt.Println(1)

}