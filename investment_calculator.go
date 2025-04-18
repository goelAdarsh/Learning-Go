package main

// func main() {
// 	// EXAMPLE 1
// 	// fmt.Print("Hello World!");
// 	// -----------------------------

// 	// EXAMPLE 2
// 	// var investmentAmount = 1000
// 	// var expectedReturnRate = 5.5
// 	// var years = 10

// 	// var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))

// 	// fmt.Println(futureValue)
// 	// -----------------------------

// 	// EXAMPLE 3
// 	// var investmentAmount float64 = 1000
// 	// var expectedReturnRate = 5.5
// 	// var years float64 = 10

// 	// var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)

// 	// fmt.Println(futureValue)
// 	// -----------------------------

// 	// EXAMPLE 4
// 	// var investmentAmount float64 = 1000
// 	// expectedReturnRate := 5.5 // type inferred by Go; explicit type assignment not allowed
// 	// var years float64 = 10

// 	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)

// 	// fmt.Println(futureValue)
// 	// -----------------------------

// 	// EXAMPLE 5
// 	// var investmentAmount, years float64 = 1000, 10 // defining multiple variables in one line; cannot specify multiple type assignments

// 	// var investmentAmount, years = 1000, 10 // type inferred; can specify multiple values of different types in the same line

// 	// investmentAmount, years := 1000.0, 10.0
// 	// expectedReturnRate := 5.5

// 	// investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0

// 	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)

// 	// fmt.Println(futureValue)
// 	// -----------------------------

// 	//  EXAMPLE 6
// 	// const inflationRate = 6.5
// 	// investmentAmount, expectedReturnRate, years := 1000.0, 5.5, 10.0

// 	// futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
// 	// futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

// 	// fmt.Println(futureValue)
// 	// fmt.Println(futureRealValue)
// 	// -----------------------------

// 	// EXAMPLE 7
// 	const inflationRate = 6.5

// 	// var investmentAmount float64 = 1000 // will be overwritten by user input
// 	// expectedReturnRate, years := 5.5, 10.0

// 	var investmentAmount, expectedReturnRate, years float64

// 	fmt.Print("Investment Amount: ")
// 	fmt.Scan(&investmentAmount)

// 	fmt.Print("Expected Return Rate: ")
// 	fmt.Scan(&expectedReturnRate)

// 	fmt.Print("Years: ")
// 	fmt.Scan(&years)

// 	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
// 	futureRealValue := futureValue / math.Pow((1+inflationRate/100), years)

// 	// fmt.Println("Future Value:", futureValue) // outputs in the same line separated by a blank
// 	// fmt.Println("Future Value (adjusted after Inflation):", futureRealValue)

// 	// fmt.Printf("Future Value: %v\n", futureValue) // %v - value of a variable/constant

// 	fmt.Printf("Future Value: %v\nFuture Value (adjusted after Inflation): %v\n", futureValue, futureRealValue)     // 1708.1444583535929, 909.9730253950707

// 	fmt.Printf("Future Value: %f\nFuture Value (adjusted after Inflation): %f\n", futureValue, futureRealValue)     // 1708.144458, 909.973025
// 	fmt.Printf("Future Value: %.2f\nFuture Value (adjusted after Inflation): %.2f\n", futureValue, futureRealValue) // 1708.14, 909.97
// 	fmt.Printf("Future Value: %.0f\nFuture Value (adjusted after Inflation): %.0f", futureValue, futureRealValue)   // 1708, 910;
// }
