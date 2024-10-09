package main

import "fmt"

func main() {
	var hoursPerWeek, hourlyRate, overtimeHours, normalHours float64
	var weeklySalary, overtimePay, monthlySalary float64

	fmt.Print("Masukkan jumlah jam kerja dalam seminggu: ")
	fmt.Scanln(&hoursPerWeek)

	fmt.Print("Masukkan upah per jam: ")
	fmt.Scanln(&hourlyRate)

	if hoursPerWeek > 40 {
		normalHours = 40
		overtimeHours = hoursPerWeek - 40
		overtimePay = overtimeHours * hourlyRate * 1.5
	} else {
		normalHours = hoursPerWeek
		overtimePay = 0
	}

	weeklySalary = normalHours * hourlyRate

	monthlySalary = (weeklySalary + overtimePay) * 4

	fmt.Printf("Total gaji bulanan: Rp %.2f\n", monthlySalary)
}
