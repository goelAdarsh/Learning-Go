package main

// import "fmt"

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
// 	revenue := getUserInput("Revenue: ")
// 	expenses := getUserInput("Expenses: ")
// 	taxRate := getUserInput("Tax Rate: ")

// 	// ---------------------------------------------
// 	// earningsBeforeTax := revenue - expenses
// 	// // earningsBeforeTax := float64(revenue) - expenses
// 	// profit := earningsBeforeTax * (1 - taxRate/100)
// 	// ratio := earningsBeforeTax / profit
// 	earningsBeforeTax, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

// 	// ----OUTPUT----

// 	// fmt.Println("Earnings before Tax", earningsBeforeTax)
// 	// fmt.Println("Profit", profit)
// 	// fmt.Println("Ratio", ratio)

// 	fmt.Printf("Earnings before Tax (EBT): %.2f\n", earningsBeforeTax)
// 	fmt.Printf("Profit: %.2f\n", profit)
// 	fmt.Printf("Ratio of EBT to Profit: %.4f", ratio)
// }

// func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
// 	earningsBeforeTax := revenue - expenses
// 	profit := earningsBeforeTax * (1 - taxRate/100)
// 	ratio := earningsBeforeTax / profit
// 	return earningsBeforeTax, profit, ratio
// }

// func getUserInput(infoText string) float64 {
// 	var userInput float64

// 	fmt.Print(infoText)
// 	fmt.Scan(&userInput)

// 	return userInput
// }
