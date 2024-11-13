/*
#Soal nomor 5

Seorang pedagang sedang menghitung harga jual suatu barang yang akan dijualnya.
Pedagang tersebut menetapkan bahwa keuntungan setiap barang yang dijual adalah 5% dari modal barangnya.
Masukan terdiri dari tiga bilangan bulat positif yang menyatakan harga beli tiga barang yang akan dijual.
Keluaran berupa tiga bilangan yang menyatakan harga jual dari masing-masing barang dengan keuntungan 5%.

#Pseucode
program Suhu

kamus

	harga: integer

algoritma

	input(harga pertama, harga kedua, harga ketiga)
	keuntungan := harga : 100 * 5
	hasil akhir:= harga + keuntungan
	output(hasil)

endprogram
*/
package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Masukan harga barang: ")
	fmt.Scan(&a, &b, &c)

	fmt.Println(a/100*5 + a)
	fmt.Println(b/100*5 + b)
	fmt.Println(c/100*5 + c)
}