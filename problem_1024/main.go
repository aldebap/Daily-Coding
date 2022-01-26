////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Jan-25-2022  -  aldebap
//
//	Daily Coding problem #1024
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
)

////////////////////////////////////////////////////////////////////////////////
//	Solution entry point
////////////////////////////////////////////////////////////////////////////////

func main() {

	//	splash screen
	fmt.Printf(">>> Problem #1024\n\n")

	//	a few test cases
	type testCase struct {
		input  uint32
		result uint32
	}

	var testScenarios []testCase = []testCase{
		{input: 0xf0f0f0f0, result: 0x0f0f0f0f},
		{input: 0xcc3333cc, result: 0x33cccc33},
	}

	for _, test := range testScenarios {

		fmt.Printf("input: %032b/%08x; expected: %032b/%08x; result: %032b/%08x\n", test.input, test.input,
			test.result, test.result, reverseBits(test.input), reverseBits(test.input))
	}
}

//	reverse the bits of a uint64 number
func reverseBits(num uint32) uint32 {

	var result uint32 = 0
	var currentBit uint32 = 1

	for i := 1; i <= 32; i++ {
		if num&currentBit == 0 {
			result |= currentBit
		}
		currentBit <<= 1
	}

	return result
}
