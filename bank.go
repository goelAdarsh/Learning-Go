package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	// contains utility functions to work eith strings
)

const accountBalanceFile = "balance.txt"

func getFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	// data, _ := os.ReadFile(accountBalanceFile) // _ eplicitly tells Go that we know there's a value but we don't wanna use it
	// if the file doesn't exist, Go doesn't crashes the application unlike other programming languages, instead returns an empty []byte collection, which is then converted into an empty string (below), which is then converted to 0
	// all functions in Go, are and should be written in such a way that it doesn't crashes the application

	data, err := os.ReadFile(fileName)

	// nil - special value in Go that indicates absence of a useful value
	if err != nil {
		// return 1000, errors.New("Failed to read balance from the file.") // default balance
		return defaultValue, errors.New("Failed to read data from the file.") // default balance
	}

	valueText := string(data) // []byte --> string

	value, err := strconv.ParseFloat(valueText, 64) // string --> float
	if err != nil {
		// return 1000, errors.New("Failed to parse stored balance value.") // default balance
		return defaultValue, errors.New("Failed to parse the stored value.") // default balance

	}

	return value, nil // nil states no errors
}

func writeFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value) // to convert balance from float64 to a string, which can be converted into []byte
	os.WriteFile(fileName, []byte(valueText), 0644)
	// 0644 : 6-Owner(rw-); 4-Group(r--); 4-Others(r--)
}

func main() {
	var accountBalance, err = getFloatFromFile(accountBalanceFile, 1000)

	if err != nil {
		fmt.Println("ERROR 💥")
		fmt.Println(err)
		fmt.Println("----------")
		// return // if an error occurs, exit the application
		panic("Sorry! Can't continue.") // alternative to return but makes it look more like a crash
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
			writeFloatToFile(accountBalance, accountBalanceFile)
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
			writeFloatToFile(accountBalance, accountBalanceFile)
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
