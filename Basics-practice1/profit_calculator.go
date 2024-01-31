package main

import "fmt"

func main() {
  revenue :=	getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	earningBeforeTax, earningAfterTax, ratio := calculateFinalcials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", earningBeforeTax)
	fmt.Printf("%.1f\n", earningAfterTax)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinalcials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	earningAfterTax := earningBeforeTax * (1 - taxRate/100)
	ratio := earningBeforeTax / earningAfterTax
	return earningBeforeTax, earningAfterTax, ratio
}


func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	return userInput
}