package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalnaceFile = "blalance.txt"

func getBalanceFromFile() float64 {
	data, _ :=	os.ReadFile(accountBalnaceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalnaceFile, []byte(balanceText), 0644)
} 

func main() {
	var accountBalnace float64 = getBalanceFromFile()

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
	
		switch choice {
			case 1:
				fmt.Println("Your balance is:", accountBalnace)
			case 2:
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
				writeBalanceToFile(accountBalnace)
			case 3:
				fmt.Print("Your withdraw: ")
				var withdrawAmount float64
				fmt.Scan(&withdrawAmount)
		
				if withdrawAmount <= 0 {
					fmt.Println("Invalid amount! Must be greater than 0")
					continue
				}
		
				if withdrawAmount > accountBalnace {
					fmt.Println("Invalid amount! You can't withdraw more than you have.")
					continue
				}
		
				accountBalnace -= withdrawAmount
				fmt.Println("Balance updated! New amount:", accountBalnace)
				writeBalanceToFile(accountBalnace)
			default:
				fmt.Println("Good bye!")
				fmt.Println("Thanks for choosing our Bank!")
				return
				//break
		}
	}	
}