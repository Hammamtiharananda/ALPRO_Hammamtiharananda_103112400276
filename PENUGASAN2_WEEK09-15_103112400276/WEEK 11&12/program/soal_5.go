package main

import "fmt"

func main() {

	var totalBelanja, hargaAkhir int
	var kartu, diskon, cashback bool

	fmt.Scan(&totalBelanja, &kartu)

	if totalBelanja >= 200000 && kartu == true {
		diskon = true
		cashback = true
	} else if totalBelanja >= 100000 {
		diskon = true
		cashback = false
	}

	if diskon == true && cashback == true {
		hargaAkhir = (totalBelanja - (totalBelanja * 10 / 100)) - 75000
	} else if diskon == true && cashback == false {
		hargaAkhir = totalBelanja - (totalBelanja * 10 / 100)
	} else {
		hargaAkhir = totalBelanja
	}

	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Println("Rp", hargaAkhir)

}
