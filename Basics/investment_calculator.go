package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate float64 = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5 

	fmt.Print("Enter investment amount: ")
	// pointer
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter years: ")
	fmt.Scan(&years)


	  
	futureValue := investmentAmount * math.Pow((1 + expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow((1 + inflationRate/100), years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}