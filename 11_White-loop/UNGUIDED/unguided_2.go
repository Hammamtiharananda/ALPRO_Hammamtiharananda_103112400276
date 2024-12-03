package main

import "fmt"

func main() {

	var bilangan, digit int

	fmt.Print("Masukan bilangan: ")
	fmt.Scan(&bilangan)

	// Perulangan ketika bilangan > 0 maka perulangan akan terus berjalan, ketika sudah tidak > 0 maka perulangan akan berhenti
	for bilangan > 0 {
		digit = bilangan % 10
		fmt.Println(digit)
		bilangan /= 10
	}

	fmt.Println("Bilangan yang dimasukan harus > 0")
}