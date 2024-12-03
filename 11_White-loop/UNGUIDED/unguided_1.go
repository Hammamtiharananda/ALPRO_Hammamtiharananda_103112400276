package main

import "fmt"

func main() {

	var akun, passoword, input string
	var jml int

	akun = "Admin"
	passoword = "Admin"

	fmt.Print("Masukan Akun & Password: ")
	fmt.Scan(&input, &input)

	// Perulangan yang akan berjalan ketika akun dan password tidak sesuai
	for input != akun && input != passoword {

		// Perulangan akan menampilkan inputan untuk memasukan password terus sampai akun dan password benar
		fmt.Print("Salah, Masukan Akun & Password: ")
		fmt.Scan(&input, &input)
		// Variabel jml disini digunakan untuk menyimpan data jumlah perulangan yang berjalan ada berapa kali.
		jml++

	}

	fmt.Println("Total Salah Memasukan Password Adalah:", jml)

}