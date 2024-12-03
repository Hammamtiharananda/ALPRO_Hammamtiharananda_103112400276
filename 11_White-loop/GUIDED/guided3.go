package main

import "fmt"

func main() {

	// Membuat variabel
	var n, s1, s2, j, temp int

	// Inputan untuk memasukan bilangan
	fmt.Print("Masukan bilangan fibonacci: ")
	fmt.Scan(&n)

	// Nilai default untuk s1 s2 dan j
	s1 = 0
	s2 = 1
	j = 0

	// Perulangan untuk menentukan bilangan fibonacci
	for j < n {
		fmt.Print(s1, " ")
		temp = s1 + s2
		s1 = s2
		s2 = temp
		j++
	}

	fmt.Println()

}