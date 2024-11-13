/*
Soal nomer 8 Faktorial
Buatlah algoritma untuk menghitung nilai faktorial dari suatu bilangan bulat positif. Masukan terdiri dari sebuah bilangan bulat positif n.
Keluaran berupa sebuah bilangan bulat yang menyatakan factorial dari n.


#Pseucode

program bilangan
kamus
    n : integer
	fakrotial integer dengan value 1

algoritma
    input(n / bilangan positif)
    percabangan
		jika bilangan < 0
			Eksekusi output(Bilangan harus positif)

		perulangan i = 1 dan i kurang dari inputan
			rumus = bilangan terus di kali dengan rumus *=

    output (faktorial)

endprogram
*/

package main

import "fmt"

func main() {
	var n int
	var faktorial int = 1

	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	if n < 0 {
		fmt.Println("Bilangan harus positif!")
		return
	}

	for i := 1; i <= n; i++ {
		faktorial *= i
	}

	fmt.Printf("%d = %d\n", n, faktorial)
}