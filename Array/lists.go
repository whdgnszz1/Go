package main

import "fmt"

func main () {
	var productNames [4]string = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	featuredPrices := prices[1:]
	featuredPrices[0] = 199.99
	highlitedPrices := featuredPrices[:1]
	fmt.Println(highlitedPrices)
	fmt.Println(prices)
	fmt.Println(len(featuredPrices), cap(featuredPrices))
	fmt.Println(len(highlitedPrices), cap(highlitedPrices))

	highlitedPrices = highlitedPrices[:3]
	fmt.Println(highlitedPrices)
	fmt.Println(len(highlitedPrices), cap(highlitedPrices))
}

