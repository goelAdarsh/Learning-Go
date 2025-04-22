package main

import (
	"fmt"

	"example.com/investment-calculator/fileOperations" // importing own custom package
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileOperations.GetFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Println("ERROR 💥")
		fmt.Println(err)
		fmt.Println("----------")
		// return // if an error occurs, exit the application
		// panic("Sorry! Can't continue.") // alternative to return but makes it look more like a crash
	}

	fmt.Println("|| Welcome to Go Bank ||")

	for {
		presentOptions()

		fmt.Print("Your choice: ")
		var choice int
		fmt.Scan(&choice) // if a value other than int is entered, it will be assigned a default value of 0

		switch choice {
		case 1:
			fmt.Println("Your balance is", accountBalance)

		case 2:
			fmt.Print("Your deposit amount: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			if depositAmount <= 0 {
				fmt.Println("Oops! Invalid amount. Must be greater than zero.")
				continue
			}

			accountBalance += depositAmount
			fileOperations.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount:", accountBalance)

		case 3:
			fmt.Print("Your withdrawal amount: ")
			var withdrawalAmount float64
			fmt.Scan(&withdrawalAmount)

			if withdrawalAmount <= 0 {
				fmt.Println("Oops! Invalid amount. Must be greater than zero.")
				continue
			}

			if withdrawalAmount > accountBalance {
				fmt.Println("Oops! You are trying to withdraw an amount greater than your current account balance.")
				continue
			}

			accountBalance -= withdrawalAmount
			fileOperations.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount:", accountBalance)
		default:
			fmt.Println("Thank You for choosing Go Bank.")
			fmt.Println("Please visit us again! Goodbye 👋")
			// break // cannot use; it breaks out of switch-case and not of for-loop; hence, better solution --> if-else-if-else
			return
		}
		// end of switch
	}
	// end of for loop
}
