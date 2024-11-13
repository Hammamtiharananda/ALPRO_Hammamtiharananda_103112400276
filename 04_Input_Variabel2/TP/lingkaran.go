package main

import (
	"fmt"
)

func main() {
	var jari float32
	var luas float32
	var keliling float32

	fmt.Println("Masukkan jari-jari: ")
	fmt.Scanln(&jari)
	luas = 3.14 * jari * jari
	keliling = 2 * 3.14 * jari
	fmt.Println("jadi hasil luas lingkarannya adalah: ", luas)
	fmt.Println("jadi hasil keliling lingkarannya adalah: ", keliling)
}