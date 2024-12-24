package main
import "fmt"

func main () {
	var bilangan int
	fmt.Print("masukan bilangan positif: ")
	fmt.Scanln(&bilangan)

	fmt.Print("faktor dari",bilangan,"adalah: ")
	for i := 1; i <= bilangan; i++{
		if bilangan %i == 0 {
			fmt.Print(i," ")
		}
	}
}