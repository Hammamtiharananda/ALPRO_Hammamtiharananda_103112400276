
package main

import "fmt"

func main() {

	var harga, total int

	fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
	fmt.Scan(&harga)

	// Perulangan untuk menghitung total belanja dan akan berhenti ketika memasukkan angka 0
	for harga != 0 {
		total += harga
		fmt.Print("Masukkan harga barang (ketik 0 untuk selesai): ")
		fmt.Scan(&harga)
	}

	fmt.Println("Total Belanja Anda:", total)

}
