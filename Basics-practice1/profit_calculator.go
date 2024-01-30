package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Enter revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Enter tax rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - taxRate/100)
	ratio := earningAfterTax / revenue * 100

	fmt.Println("EBT: ", earningBeforeTax)
	fmt.Println("Profit: ", earningAfterTax)
	fmt.Println("Profit Ratio: ", ratio)
}