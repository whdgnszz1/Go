package main

import "fmt"

func main () {
	prices := []float64{10.99, 8.99}
	fmt.Println(prices[0:1])
	prices[1] = 9.99
	// prices[2] = 11.99

	// fmt.Println(prices)
	updatedPrices := append(prices, 5.99)
	fmt.Println(updatedPrices)

	prices = append(prices, 5.99)
	prices = prices[1:]
	fmt.Println(prices)
}

// func main () {
// 	var productNames [4]string = [4]string{"A Book"}
// 	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
// 	fmt.Println(prices)

// 	productNames[2] = "A Carpet"

// 	fmt.Println(productNames)

// 	fmt.Println(prices[2])

// 	featuredPrices := prices[1:]
// 	featuredPrices[0] = 199.99
// 	highlitedPrices := featuredPrices[:1]
// 	fmt.Println(highlitedPrices)
// 	fmt.Println(prices)
// 	fmt.Println(len(featuredPrices), cap(featuredPrices))
// 	fmt.Println(len(highlitedPrices), cap(highlitedPrices))

// 	highlitedPrices = highlitedPrices[:3]
// 	fmt.Println(highlitedPrices)
// 	fmt.Println(len(highlitedPrices), cap(highlitedPrices))
// }
