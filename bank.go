package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// 	"strconv"
// 	// contains utility functions to work eith strings
// )

// const accountBalanceFile = "balance.txt"

// func getBalanceFromFile() (float64, error) {
// 	// data, _ := os.ReadFile(accountBalanceFile) // _ eplicitly tells Go that we know there's a value but we don't wanna use it
// 	// if the file doesn't exist, Go doesn't crashes the application unlike other programming languages, instead returns an empty []byte collection, which is then converted into an empty string (below), which is then converted to 0
// 	// all functions in Go, are and should be written in such a way that it doesn't crashes the application

// 	data, err := os.ReadFile(accountBalanceFile)

// 	// nil - special value in Go that indicates absence of a useful value
// 	if err != nil {
// 		return 1000, errors.New("Failed to read balance from the file.") // default balance
// 	}

// 	balanceText := string(data) // []byte --> string

// 	balance, err := strconv.ParseFloat(balanceText, 64) // string --> float
// 	if err != nil {
// 		return 1000, errors.New("Failed to parse stored balance value.") // default balance
// 	}

// 	return balance, nil // nil states no errors
// }

// func writeBalanceToFile(balance float64) {
// 	balanceText := fmt.Sprint(balance) // to convert balance from float64 to a string, which can be converted into []byte
// 	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
// 	// 0644 : 6-Owner(rw-); 4-Group(r--); 4-Others(r--)
// }

// func main() {
// 	var accountBalance, err = getBalanceFromFile()

// 	if err != nil {
// 		fmt.Println("ERROR 💥")
// 		fmt.Println(err)
// 		fmt.Println("----------")
// 		// return // if an error occurs, exit the application
// 		panic("Sorry! Can't continue.") // alternative to return but makes it look more like a crash
// 	}

// 	fmt.Println("|| Welcome to Go Bank ||")

// 	for {
// 		fmt.Println("What do you want to do?")
// 		fmt.Println("1. Check balance")
// 		fmt.Println("2. Deposit money")
// 		fmt.Println("3. Withdraw money")
// 		fmt.Println("4. Exit")

// 		fmt.Print("Your choice: ")
// 		var choice int
// 		fmt.Scan(&choice) // if a value other than int is entered, it will be assigned a default value of 0

// 		switch choice {
// 		case 1:
// 			fmt.Println("Your balance is", accountBalance)

// 		case 2:
// 			fmt.Print("Your deposit amount: ")
// 			var depositAmount float64
// 			fmt.Scan(&depositAmount)

// 			if depositAmount <= 0 {
// 				fmt.Println("Oops! Invalid amount. Must be greater than zero.")
// 				continue
// 			}

// 			accountBalance += depositAmount
// 			writeBalanceToFile(accountBalance)
// 			fmt.Println("Balance updated! New amount:", accountBalance)

// 		case 3:
// 			fmt.Print("Your withdrawal amount: ")
// 			var withdrawalAmount float64
// 			fmt.Scan(&withdrawalAmount)

// 			if withdrawalAmount <= 0 {
// 				fmt.Println("Oops! Invalid amount. Must be greater than zero.")
// 				continue
// 			}

// 			if withdrawalAmount > accountBalance {
// 				fmt.Println("Oops! You are trying to withdraw an amount greater than your current account balance.")
// 				continue
// 			}

// 			accountBalance -= withdrawalAmount
// 			writeBalanceToFile(accountBalance)
// 			fmt.Println("Balance updated! New amount:", accountBalance)
// 		default:
// 			fmt.Println("Thank You for choosing Go Bank.")
// 			fmt.Println("Please visit us again! Goodbye 👋")
// 			// break // cannot use; it breaks out of switch-case and not of for-loop; hence, better solution --> if-else-if-else
// 			return
// 		}
// 		// end of switch

// 	}
// 	// end of for loop

// }
