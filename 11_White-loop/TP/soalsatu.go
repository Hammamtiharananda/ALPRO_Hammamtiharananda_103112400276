package main

import "fmt"

func main() {
	var input string
	var max int

	password := "qwrty"
	max = 3

	for i := 0; i < max; {
		fmt.Print("Masukkan Password yang benar: ")
		fmt.Scan(&input)
		if input == password {
			fmt.Print("Anda Berhasil Masuk")
			return
		} else {
			fmt.Print("Password Anda Salah\n")
			i++
		}

		if i == max {
			print("Anda Melebihi Limit")
		}
	}
}