package main

import "fmt"

func main() {
	// Membuat variabel
	var token string

	// Inputan untuk memasukan password
	fmt.Print("Masukan Token: ")
	fmt.Scan(&token)

	// Perulangan untuk mengecek apakah password yang dimasukan benar atau salah
	for token != "12345abcde" {
		fmt.Println("Token Salah, coba lagi!")

		fmt.Print("Masukan Password: ")
		fmt.Scan(&token)
	}

	// Tampil jika password benar
	fmt.Println("Selamat, anda berhasil login!")
}