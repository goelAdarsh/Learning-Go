package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("|| Welcome to Go Bank ||")

	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

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

			fmt.Println("Balance updated! New amount:", accountBalance)
		default:
			fmt.Println("Thank You for choosing Go Bank.")
			fmt.Println("Please visit us again! Goodbye 👋")
			return
		}
		// end of switch

	}
	// end of for loop

}
