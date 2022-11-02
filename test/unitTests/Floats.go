/*
	This unit test explores using math.Float__bits and math.Float__frombits,
	where the __ is 32, but can be changed to reflect a 64-bit implementation
*/

package main

import (
	"codeBytes/pkg"
	"fmt"
	"math"
	"strconv"
)

func main() {
	// The value being tested is inp
	inp := "3.1415"
	fmt.Println("Input is:", inp)

	// These store final and intermediate results
	var result []byte
	goBack := uint32(0)

	// Checks that inp is a float
	f, err := strconv.ParseFloat(inp, 32)
	pkg.HandleError(err)

	// Converts float to binary string via IEEE standard
	bin := fmt.Sprintf("%032b", math.Float32bits(float32(f)))

	// Parses binary in bytes and populates result
	for i := uint(0); i < 32; i += 8 {
		singleBin, _ := strconv.ParseUint(bin[i:i+8], 2, 64)
		result = append(result, byte(singleBin))
	}

	fmt.Println("Byte array is:", result)

	// Converts byte result back to its uint32 representation
	for i := 3; i > -1; i-- {
		goBack += uint32(result[3-i]) * uint32(pkg.Exp(2, 8*i))
	}

	fmt.Println("Check for converting back:", math.Float32frombits(goBack))
}
