
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


package main

import "fmt"

func main() {
    var celsius float64

    fmt.Print("Masukkan suhu dalam derajat Celsius: ")
    fmt.Scan(&celsius)

    fahrenheit := (celsius * 9 / 5) + 32
    reamur := celsius * 4 / 5
    kelvin := celsius + 273.15

    fmt.Printf("Derajat Fahrenheit: %.2f\n", fahrenheit)
    fmt.Printf("Derajat Reamur: %.2f\n", reamur)
    fmt.Printf("Derajat Kelvin: %.2f\n", kelvin)
}

package main

import (
    "fmt"
    "math"
)

func main() {
    var r float64
    const pi = 3.1415926535

    fmt.Print("Masukkan jari-jari bola: ")
    fmt.Scan(&r)

    volume := (4.0 / 3.0) * pi * math.Pow(r, 3)
    luas := 4 * pi * math.Pow(r, 2)

    fmt.Printf("Volume bola = %.4f\n", volume)
    fmt.Printf("Luas kulit bola = %.4f\n", luas)
}



package main

import "fmt"

func main() {
    var tahun int

    fmt.Print("Masukkan tahun: ")
    fmt.Scan(&tahun)

    if (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0) {
        fmt.Println("Kabisat: true")
    } else {
        fmt.Println("Kabisat: false")
    }
}



