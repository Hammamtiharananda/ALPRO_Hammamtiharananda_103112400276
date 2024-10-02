
package main

import "fmt"

func main() {
    var fx, x float64

    fmt.Print("Masukkan nilai f(x): ")
    fmt.Scan(&fx)

    // Menghitung nilai x dari persamaan f(x) = 2/(x+5) + 5
    x = (2 / (fx - 5)) - 5

    fmt.Printf("Nilai x = %.3f\n", x)
}