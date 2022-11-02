/*
	Takes user input, infers type, and converts it to bytes accordingly
*/

package main

import (
	"codeBytes/pkg"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Appends the result by parsing a binary string into bytes
func createResult(result *[]byte, binary string, size uint) {
	for i := uint(0); i < size/8; i++ {
		singleBin, _ := strconv.ParseUint(binary[8*i:8*(i+1)], 2, 64)
		*result = append(*result, byte(singleBin))
	}
}

// Encodes an unsigned integer into bytes
func encodeUint(value int64, size *uint, result *[]byte) {
	for value >= pkg.Exp(2, int(*size)) {
		*size += 8
	}

	bin := fmt.Sprintf("%0"+strconv.FormatUint(uint64(*size), 10)+"b", value)

	createResult(result, bin, *size)
}

func main() {
	userEntry := pkg.Input("Enter a primitive value and I will convert it to a byte array\nbased on its type: ")

	// Result variables
	var result []byte
	var typ string
	size := uint(8)

	// Handles Booleans
	switch strings.ToLower(userEntry) {
	case "true":
		typ = "Boolean"
		result = append(result, 1)
	case "false":
		typ = "Boolean"
		result = append(result, 0)
	}

	// Ensures the hex and bin conversions don't go out of indexing range
	safeEntry := "safe"
	if len(userEntry) > 2 {
		safeEntry = userEntry[2:]
	}

	// Handles all the string to <type> conversions
	binVal, binErr := strconv.ParseUint(safeEntry, 2, 64)
	hexVal, hexErr := strconv.ParseUint(safeEntry, 16, 64)
	uintVal, uintErr := strconv.ParseUint(userEntry, 10, 64)
	intVal, intErr := strconv.ParseInt(userEntry, 10, 64)
	float32Val, float32Err := strconv.ParseFloat(userEntry, 32)
	float64Val, float64Err := strconv.ParseFloat(userEntry, 64)

	// This if...else block determines type based on a hierarchy of error checking
	if binErr == nil && userEntry[:2] == "0b" {
		typ = "Binary"
		encodeUint(int64(binVal), &size, &result)
	} else if hexErr == nil && userEntry[:2] == "0x" {
		typ = "Hexadecimal"
		encodeUint(int64(hexVal), &size, &result)
	} else if uintErr == nil {
		typ = "Unsigned Integer"
		encodeUint(int64(uintVal), &size, &result)
	} else if intErr == nil {
		typ = "Signed Integer"

		// Controls the sign bit
		signed := "0"
		if intVal < 0 {
			signed = "1"
			intVal *= -1
		}

		// Determines the number of bytes needed to fit the integer
		for intVal >= pkg.Exp(2, int(size-1)) {
			size += 8
		}

		// Converts integer value and the sign to binary
		bin := fmt.Sprintf("%s%0"+strconv.FormatUint(uint64(size-1), 10)+"b", signed, intVal)

		createResult(&result, bin, size)
	} else if float32Err == nil {
		typ = "Floating-point 32-bit"
		size = 32
		bin := fmt.Sprintf("%032b", math.Float32bits(float32(float32Val)))

		createResult(&result, bin, size)
	} else if float64Err == nil {
		typ = "Floating-point 64-bit"
		size = 64
		bin := fmt.Sprintf("%064b", math.Float64bits(float64Val))

		createResult(&result, bin, size)
	} else {
		typ = "String"

		// Converts each character to its ASCII byte value
		for i := range userEntry {
			result = append(result, userEntry[i])
		}
	}

	// Prints results
	fmt.Println("Type is:", typ)
	fmt.Println("Byte array is:", result)
}
