package main

import "fmt"

func main() {
	var accountBalnace float64 = 1000

	fmt.Println("Welcome to Go Bank!")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")
	
		var choice int
		fmt.Print("Your Choice: ")
		fmt.Scan(&choice)
	
		// wantsCheckBalance := choice == 1
	
		if choice == 1 {
			fmt.Println("Your balance is:", accountBalnace)
		} else if choice == 2 {
			fmt.Print("Your deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
	
			if depositAmount <= 0 {
				fmt.Println("Invalid amount! Must be greater than 0")
				// return
				continue
			}
	
			accountBalnace += depositAmount
			fmt.Println("Balance updated! New amount:", accountBalnace)
		} else if choice == 3 {
			fmt.Print("Your withdraw: ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)
	
			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount! Must be greater than 0")
				return
			}
	
			if withdrawAmount > accountBalnace {
				fmt.Println("Invalid amount! You can't withdraw more than you have.")
				return
			}
	
			accountBalnace -= withdrawAmount
			fmt.Println("Balance updated! New amount:", accountBalnace)
		} else  {
			fmt.Println("Good bye!")
			// return
			break
		} 
	}

	fmt.Println("Thanks for choosing our Bank!")
}