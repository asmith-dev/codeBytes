/*
	This script includes a variety of useful functions to condense workflow.
*/

package pkg

import (
	"fmt"
)

// HandleError simplifies general error handling.
func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

// Input is a Python-like implementation for getting user input in a condensed format.
func Input(str string) string {
	var response string
	fmt.Print(str)

	_, err := fmt.Scanln(&response)
	HandleError(err)

	return response
}

// Exp calculates integer exponents of integers
func Exp(base int64, exp int) int64 {
	result := int64(1)

	for exp > 0 {
		result *= base
		exp--
	}

	return result
}
