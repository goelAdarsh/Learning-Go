package fileOperations

// name of the package should be same as the name of the sub-folder
// file name does not need to be the same

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	// contains utility functions to work eith strings
)

// functions, variables or constants that starts with uppercase character will only be exported and hence, available in other packages
func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	// func getFloatFromFile(fileName string, defaultValue float64) (float64, error) {
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

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value) // to convert balance from float64 to a string, which can be converted into []byte
	os.WriteFile(fileName, []byte(valueText), 0644)
	// 0644 : 6-Owner(rw-); 4-Group(r--); 4-Others(r--)
}
