package main

import "fmt"

func main() {
	var umur int
	var kk string

	fmt.Print("Masukan umur anda: ")
	fmt.Scan(&umur)
	fmt.Print("Apakah sudah mempunyai KK: ya/tidak ")
	fmt.Scan(&kk)

	if umur >= 17 && kk == "ya" {

		fmt.Println("Anda bisa membuat ktp")
		fmt.Println("True")

	} else {

		fmt.Println("Anda tidak bisa membuat KTP")
		fmt.Println("False")

	}
}