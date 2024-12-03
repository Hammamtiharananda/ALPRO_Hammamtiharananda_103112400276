package main

import "fmt"

func main() {

	var x, y, jumlah int

	fmt.Print("Masukan bilangan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukan bilangan y: ")
	fmt.Scan(&y)

	// Perulangan akan berjalan jika nilai x lebih besar atau sama dengan y, kemudian ada beberapa operasi yang akan dilakukan untuk menentukan division
	for x >= y {
		x = x - y
		jumlah++
	}

	fmt.Print(jumlah)

}