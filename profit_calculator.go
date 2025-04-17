package main

import (
	"fmt"
)

func main() {
	var revenue, expenses, taxRate float64
	// var revenue int
	// var expenses, taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	// earningsBeforeTax := float64(revenue) - expenses
	profit := earningsBeforeTax * (1 - taxRate/100)
	ratio := earningsBeforeTax / profit

	fmt.Println(earningsBeforeTax)
	fmt.Println(profit)
	fmt.Println(ratio)
}
