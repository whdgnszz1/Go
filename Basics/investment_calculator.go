package main

import (
	"fmt"
	"math"
)

const inflationRate float64 = 2.5


func main() {
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5  
	 
	// fmt.Print("Enter investment amount: ")
	outputText("Enter investment amount: ")
	// pointer
	fmt.Scan(&investmentAmount)

	// fmt.Print("Enter Expected Return Rate: ")
	outputText("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("Enter years: ")
	outputText("Enter years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow((1 + expectedReturnRate/100), years)
	// futureRealValue := futureValue / math.Pow((1 + inflationRate/100), years)

	formmatedFV := fmt.Sprintf("Future Value: %.1f", futureValue)
	formmatedRFV := fmt.Sprintf("Real Future Value: %.1f", futureRealValue)
	// outputs information
	// fmt.Println("Future Value": ,futureValue)
	// fmt.Printf(`Future Value: %.1f\n

	// Future Real Value: %.1f`, futureValue, futureRealValue)
	// fmt.Println("Future Real Value:", futureRealValue)
	fmt.Println(formmatedFV, formmatedRFV)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount , expectedReturnRate , years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1 + expectedReturnRate/100), years)
	rfv = fv / math.Pow((1 + inflationRate/100), years)
	return fv ,rfv
	// return
} 