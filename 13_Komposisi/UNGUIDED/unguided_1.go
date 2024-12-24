package main

import "fmt"

func main() {

	var angka, jumlah int

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// Perulangan untuk menampilkan angka dari 1 hingga angka yang dimasukan
	for i := 1; i <= angka; i++ {
		// Percabangan untuk memeriksa apakah angka tersebut ganjil atau genap
		if i%2 != 0 {
			jumlah++
		}
	}

	fmt.Println("Terdapat", jumlah, "bilangan ganjil")
}