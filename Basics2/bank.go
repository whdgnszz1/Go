package main

import (
	"fmt"

	"example.com/bank/fileOps"
	"github.com/Pallinder/go-randomdata"
)

const accountBalnaceFile = "blalance.txt"

func main() {
	var accountBalnace, err = fileOps.GetFloatFromFile(accountBalnaceFile, 1000)
	
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("----------------")
		panic("Can't continue without balance")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()
	
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
				fileOps.WriteFloatToFile(accountBalnaceFile ,accountBalnace)
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
				fileOps.WriteFloatToFile(accountBalnaceFile ,accountBalnace)
			default:
				fmt.Println("Good bye!")
				fmt.Println("Thanks for choosing our Bank!")
				return
				//break
		}
	}	
}

