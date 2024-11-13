/*
Soal nomer 3 BMI
Sebuah program digunakan untuk menyimpan data BMI (Body Mass Index) seseorang yang berisi nama, berat badan, tinggi badan, dan BMI orang tersebut.
Masukan: terdiri dari satu string dan dua bilangan riil yang menyatakan nama, berat badan (dalam kg), dan tinggi badan (dalam meter).
14 Keluaran: terdiri dari satu bilangan riil yang menyatakan nilai BMI.
Latihan Soal Catatan: Gunakan tipe bentukan untuk menyimpan data seseorang tersebut.


#Pseucode

program bilangan
kamus
    nama : string
	berat, tinggi, hasil : float

algoritma
    input(nama)
    input(berat)
    input(tinggi)
    hasil = berat : Tinggi badan ^2
    output (hasil)

endprogram
*/

package main

import "fmt"

func main() {

	var (
		nama                           string
		beratBadan, tinggiBadan, hasil float32
	)

	fmt.Print("Masukan nama: ")
	fmt.Scan(&nama)

	fmt.Print("Masukan berat badan: ")
	fmt.Scan(&beratBadan)

	fmt.Print("Masukan tinggi badan: ")
	fmt.Scan(&tinggiBadan)

	hasil = beratBadan / (tinggiBadan * tinggiBadan)

	fmt.Print("Informasi BMI: ")
	fmt.Println("Nama: ", nama)
	fmt.Printf("Berat: %.2f\n", beratBadan)
	fmt.Printf("Tinggi: %.2f\n", tinggiBadan)
	fmt.Printf("BMI: %.2f\n", hasil)
}