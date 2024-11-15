package main

import "fmt"

func main() {
	var masukkan int

	fmt.Scan(&masukkan)

	a := masukkan / 1000
	b := masukkan / 100 % 10
	c := masukkan / 10 % 10
	d := masukkan % 1000 % 10

	if a > b && b > c && c > d {
		fmt.Print("Digit pada bilangan ", masukkan, " terutur kebawah")

	} else if a < b && b < c && c < d {
		fmt.Print("Digit pada bilangan", masukkan, "terutur keatas")
	} else {
		fmt.Print("Digit pada bilangan ", masukkan, " tidak terurut")
	}
}