package main

import "fmt"

func main() {

	var max, donatur, total, totalDonatur int

	fmt.Print("Masukkan target donasi: ")
	fmt.Scan(&max)

	// Perulangan untuk menghitung total donasi. Perulangan akan berhenti ketika total donasi mencapai target
	for selesai := false; !selesai; {
		fmt.Print("Masukkan jumlah donasi: ")
		fmt.Scan(&donatur)

		total += donatur

		selesai = (total >= max)

		totalDonatur++
		fmt.Println("Donatur :", totalDonatur, "Menyumbang", donatur, "total terkumpul:", total)

	}

	fmt.Println("Target tercapai! Total donasi:", total, "dari", totalDonatur, "donatur.")

}