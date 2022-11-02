/*
	Demonstrates methods to decode bytes according to a given type
*/

package main

import (
	"codeBytes/pkg"
	"fmt"
	"math"
	"strconv"
)

type Stack struct {
	value  [][]byte
	cur    []byte
	typ    string
	result string
}

// Simplifies appending to the stack
func (s *Stack) addNew(inp []byte) {
	s.value = append(s.value, inp)
}

func main() {
	s := new(Stack)

	// Additions to the stack for testing
	s.addNew([]byte{0})
	s.addNew([]byte{145, 63})
	s.addNew([]byte{46, 127, 240, 65})
	s.addNew([]byte{72, 101, 108, 108, 111})
	s.addNew([]byte{128, 92, 112, 61, 34, 101, 120, 88})

	fmt.Println("Current stack values (by index) are:")
	for i, val := range s.value {
		fmt.Println("("+strconv.FormatInt(int64(i), 10)+")   =>  ", val)
	}

	inp, err := strconv.ParseUint(pkg.Input("\nEnter the index of the value to decode: "), 10, 64)
	pkg.HandleError(err)
	s.cur = s.value[inp]

	s.typ = pkg.Input("Enter the type used to interpret the value: ")

	/*
		Valid types:

		bool: boolean
		bin: binary
		hex: hexadecimal
		uint: unsigned integer
		int: signed integer
		float: floating-point 32-bit
		double: floating-point 64-bit
		str: string
	*/
	switch s.typ {
	// All zeroes means "false" whereas any non-zeroes means "true"
	case "bool":
		s.result = "false"
		for _, val := range s.cur {
			if val != 0 {
				s.result = "true"
				break
			}
		}

	// Each byte is converted to binary and concatenated
	case "bin":
		for _, val := range s.cur {
			s.result += fmt.Sprintf("%08b", val)
		}

	// Each byte is converted to hexadecimal and concatenated
	case "hex":
		for _, val := range s.cur {
			s.result += fmt.Sprintf("%02x", val)
		}

	// Adds each byte times 2 to a power, where the last byte has a power of 0,
	// the previous byte a power of 8, and so on
	case "uint":
		numVal := uint64(0)
		for i, val := range s.cur {
			numVal += uint64(val) * uint64(pkg.Exp(2, 8*(len(s.cur)-i-1)))
		}
		s.result = strconv.FormatUint(numVal, 10)

	// Works the same as uint, except the first bit of the first byte gives the sign
	case "int":
		numVal := int64(0)

		// "adjust" prevents the first bit from being accounted in numVal
		sign := int64(1)
		adjust := byte(0)
		if s.cur[0] > 127 {
			sign = -1
			adjust = 128
		}

		for i, val := range s.cur {
			numVal += int64(val-adjust) * pkg.Exp(2, 8*(len(s.cur)-i-1))
			adjust = 0
		}
		numVal *= sign
		s.result = strconv.FormatInt(numVal, 10)

	// Only works for 4 byte values. Converts to IEEE-754 float32 value
	case "float":
		if len(s.cur) != 4 {
			panic("Cannot interpret index (" + strconv.FormatUint(inp, 10) + ") as type \"" + s.typ + "\": " +
				"byte length must be 4")
		}
		numVal := uint32(0)
		for i, val := range s.cur {
			numVal += uint32(val) * uint32(pkg.Exp(2, 8*(len(s.cur)-i-1)))
		}
		s.result = strconv.FormatFloat(float64(math.Float32frombits(numVal)), 'e', -1, 32)

	// Only works for 8 byte values. Converts to IEEE-754 float64 value
	case "double":
		if len(s.cur) != 8 {
			panic("Cannot interpret index (" + strconv.FormatUint(inp, 10) + ") as type \"" + s.typ + "\": " +
				"byte length must be 8")
		}
		numVal := uint64(0)
		for i, val := range s.cur {
			numVal += uint64(val) * uint64(pkg.Exp(2, 8*(len(s.cur)-i-1)))
		}
		s.result = strconv.FormatFloat(math.Float64frombits(numVal), 'e', -1, 64)

	// Each byte is converted to its ASCII value and concatenated
	case "str":
		for _, val := range s.cur {
			s.result += string(val)
		}

	// Catches unrecognized types
	default:
		panic("Unrecognized type \"" + s.typ + "\"")
	}

	fmt.Println("\nResult is: \"" + s.result + "\"")
}
