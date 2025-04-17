package main

import (
	"fmt"
	"math"
)

func main() {
	// EXAMPLE 1
	// fmt.Print("Hello World!");
	// -----------------------------

	// EXAMPLE 2
	// var investmentAmount = 1000
	// var expectedReturnRate = 5.5
	// var years = 10

	// var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))

	// fmt.Println(futureValue)
	// -----------------------------

	// EXAMPLE 3
	// var investmentAmount float64 = 1000
	// var expectedReturnRate = 5.5
	// var years float64 = 10

	// var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	// fmt.Println(futureValue)
	// -----------------------------

	// EXAMPLE 4
	// var investmentAmount float64 = 1000
	// expectedReturnRate := 5.5 // type inferred by Go; explicit type assignment not allowed
	// var years float64 = 10

	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	// fmt.Println(futureValue)
	// -----------------------------

	// EXAMPLE 5
	// var investmentAmount, years float64 = 1000, 10 // defining multiple variables in one line; cannot specify multiple type assignments

	// var investmentAmount, years = 1000, 10 // type inferred; can specify multiple values of different types in the same line

	// investmentAmount, years := 1000.0, 10.0
	// expectedReturnRate := 5.5

	investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	fmt.Println(futureValue)
}
