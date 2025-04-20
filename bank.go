package main

import (
	"fmt"
	"os"
	"strconv"
	// contains utility functions to work eith strings
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile) // _ eplicitly tells Go that we know there's a value but we don't wanna use it
	balanceText := string(data)                // []byte --> string
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance) // to convert balance from float64 to a string, which can be converted into []byte
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
	// 0644 : 6-Owner(rw-); 4-Group(r--); 4-Others(r--)
}

func main() {
	var accountBalance = getBalanceFromFile()

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
			writeBalanceToFile(accountBalance)
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
			writeBalanceToFile(accountBalance)
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
