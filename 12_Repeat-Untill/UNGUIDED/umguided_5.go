package main

import "fmt"

func main() {

	var kantong1, kantong2, selisih float32
	var oleng bool

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1, &kantong2)

		// jika kantong 1 ataupun kantong 2 kurang dari 0 maka program selesai
		if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Proses selesai")
			break
		}

		// jika jumlah kedua kanton lebih dari 150 maka program selesai
		if kantong1+kantong2 > 150 {
			fmt.Println("Proses selesai")
			break
		}

		// digunakan untuk menentukan selisih antara kantong
		if kantong1 > kantong2 {
			selisih = kantong1 - kantong2
		} else {
			selisih = kantong2 - kantong1
		}

		oleng = selisih >= 9
		fmt.Println("Sepeda motor pak Andi akan oleng:", oleng)
	}

}