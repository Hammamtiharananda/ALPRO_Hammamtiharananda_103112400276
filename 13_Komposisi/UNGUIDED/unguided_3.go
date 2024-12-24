package main

import "fmt"

func main() {

	var gelas1, gelas2, gelas3, gelas4 string
	var berhasil bool

	berhasil = true

	// Perulangan untuk menampilkan beberapa kali inputan kemudian diperiksa apakah true atau false
	for i := 1; i <= 5; i++ {
		fmt.Print("Percobaan ", i, ": ")
		fmt.Scan(&gelas1, &gelas2, &gelas3, &gelas4)

		if !(gelas1 == "merah" && gelas2 == "kuning" && gelas3 == "hijau" && gelas4 == "ungu") {
			berhasil = false
		}
	}

	fmt.Print("Berhasil: ", berhasil)

}