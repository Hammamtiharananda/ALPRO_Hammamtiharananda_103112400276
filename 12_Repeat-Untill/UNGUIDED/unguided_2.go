package main

import "fmt"

func main() {

	var angka, done float32

	fmt.Print("Masukkan angka: ")
	fmt.Scan(&angka)

	// done disini digunakan untuk mentriger kapan perulangan berhenti
	done = float32(int(angka)) + 1

	for selesai := false; !selesai; {

		// Nilai + 0.1 digunakan untuk menjumlahkan angka dengan nilai 0.1 terus menerus sampai target done
		angka += 0.1

		// Perhitungan ini digunakan untuk membulatkan angka
		angka = float32(int(angka*10+0.5)) / 10

		// Ketika angka >= done, maka perulangan berhenti dan tidak ada angka dibelakang koma yang ditampilkan, ketika false maka 1 angka dibelakang koma akan ditampilkan
		if angka >= done {
			fmt.Printf("%.0f\n", angka)
			selesai = true
		} else {
			fmt.Printf("%.1f\n", angka)
		}

	}

}