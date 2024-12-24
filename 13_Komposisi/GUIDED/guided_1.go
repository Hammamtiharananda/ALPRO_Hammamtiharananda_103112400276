package main

import "fmt"

func main() {
	var bilangan int
	fmt.Print("Masukan bilanganya cok= ")
	fmt.Scan(&bilangan)
	
	for i := 1; i <= bilangan; i +=1 {
		if i % 2 != 0 {
			fmt.Print(i," ")
		}
	}
}