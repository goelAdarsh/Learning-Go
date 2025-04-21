package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	// ---------------------------------------------
// 	// var revenue, expenses, taxRate float64
// 	// var revenue int
// 	// var expenses, taxRate float64

// 	// fmt.Print("Revenue: ")
// 	// fmt.Scan(&revenue)

// 	// fmt.Print("Expenses: ")
// 	// fmt.Scan(&expenses)

// 	// fmt.Print("Tax Rate: ")
// 	// fmt.Scan(&taxRate)

// 	// ---------------------------------------------
// 	// revenue, err := getUserInput("Revenue: ")
// 	// if err != nil {
// 	// 	fmt.Println("ERROR 💥")
// 	// 	fmt.Println(err)
// 	// 	fmt.Println("----------")
// 	// 	panic("Sorry! Can't continue.")
// 	// }

// 	// expenses, err := getUserInput("Expenses: ")
// 	// if err != nil {
// 	// 	fmt.Println("ERROR 💥")
// 	// 	fmt.Println(err)
// 	// 	fmt.Println("----------")
// 	// 	panic("Sorry! Can't continue.")
// 	// }

// 	// taxRate, err := getUserInput("Tax Rate: ")
// 	// err, here, refers to the latest value that was assigned to it
// 	// if err != nil {
// 	// 	fmt.Println("ERROR 💥")
// 	// 	fmt.Println(err)
// 	// 	fmt.Println("----------")
// 	// 	panic("Sorry! Can't continue.")
// 	// }

// 	revenue, err1 := getUserInput("Revenue: ")
// 	expenses, err2 := getUserInput("Expenses: ")
// 	taxRate, err3 := getUserInput("Tax Rate: ")
// 	if err1 != nil || err2 != nil || err3 != nil {
// 		fmt.Println("ERROR 💥")
// 		fmt.Println(err1) // since the error messages will always be the same, we can use err1
// 		fmt.Println("----------")
// 		panic("Sorry! Can't continue.")
// 	}

// 	// ---------------------------------------------
// 	// earningsBeforeTax := revenue - expenses
// 	// // earningsBeforeTax := float64(revenue) - expenses
// 	// profit := earningsBeforeTax * (1 - taxRate/100)
// 	// ratio := earningsBeforeTax / profit
// 	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
// 	writeResultsToFile(earningsBeforeTax, profit, ratio)

// 	// ----OUTPUT----

// 	// fmt.Println("Earnings before Tax", earningsBeforeTax)
// 	// fmt.Println("Profit", profit)
// 	// fmt.Println("Ratio", ratio)

// 	fmt.Printf("Earnings before Tax (EBT): %.2f\n", earningsBeforeTax)
// 	fmt.Printf("Profit: %.2f\n", profit)
// 	fmt.Printf("Ratio of EBT to Profit: %.4f", ratio)
// }

// func writeResultsToFile(earningsBeforeTax, profit, ratio float64) {
// 	text := fmt.Sprintf("Earnings Before Tax (EBT): %.2f\nProfit: %.2f\nRatio of EBT to Profit: %.4f", earningsBeforeTax, profit, ratio)
// 	os.WriteFile("results.txt", []byte(text), 0644)
// }

// func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
// 	earningsBeforeTax := revenue - expenses
// 	profit := earningsBeforeTax * (1 - taxRate/100)
// 	ratio := earningsBeforeTax / profit
// 	return earningsBeforeTax, profit, ratio
// }

// func getUserInput(infoText string) (float64, error) {
// 	fmt.Print(infoText)
// 	var userInput float64
// 	fmt.Scan(&userInput)

// 	if userInput <= 0 {
// 		return 0, errors.New("Value must be greater than zero.")
// 	}

// 	return userInput, nil
// }
