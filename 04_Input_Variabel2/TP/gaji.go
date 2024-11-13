package main

import (
	"fmt"
)

func main() {
	var upah, jam, gajinormal, gajilembur float64

	fmt.Println("Masukan jumlah jam kerja karyawan: ")
	fmt.Scanln(&jam)
	fmt.Println("Masukan Upah karyawan: ")
	fmt.Scanln(&upah)

	if jam > 40 {
		gajinormal = (40 * upah) + ((jam - 40) * 1.5 * upah)
		gajilembur = gajinormal * 4 //terdapat 4 minggu dalam 1 bulan

		fmt.Printf("jadi total gaji bulanan berserta lemburan adalah: Rp %.2f\n", gajilembur)

	} else {
		gajinormal = jam * upah * 4
		fmt.Println("Jadi total gaji bulanan adalah: Rp", gajinormal)
	}
}