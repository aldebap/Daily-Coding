////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Jan-25-2022  -  aldebap
//
//	Daily Coding problem #1024
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"strconv"
)

////////////////////////////////////////////////////////////////////////////////
//	Solution entry point
////////////////////////////////////////////////////////////////////////////////

func main() {

	//	splash screen
	fmt.Printf(">>> Problem #1029\n\n")

	//	a few test cases
	type testCase struct {
		input  uint32
		result uint32
	}

	var testScenarios []testCase = []testCase{
		{input: 16, result: 3},
		{input: 68, result: 7},
		{input: 87, result: 6},
	}

	for _, test := range testScenarios {

		coins := minCoinsForAmount(test.input)
		coinsStr := "[ "

		//	format a string with all the coins
		for i, coin := range coins {

			if i > 0 {
				coinsStr += ", "
			}
			coinsStr += strconv.Itoa(int(coin))
		}
		coinsStr += " ]"

		fmt.Printf("input: %d cents; expected coins: %s; result: %d\n", test.input, coinsStr,
			len(coins))
	}
}

//	return a list of minimum coins that sum to the input amount (in cents)
func minCoinsForAmount(amount uint32) []uint32 {

	//	the algorithm requires that the order is from the highest value coin to the lowest
	var availableCoins [4]uint32 = [4]uint32{25, 10, 5, 1}

	var result []uint32
	var amountLeft uint32 = amount

	for i := 0; i < len(availableCoins); i++ {

		if amountLeft <= 0 {
			break
		}

		for j := uint32(1); j <= amountLeft/availableCoins[i]; j++ {

			result = append(result, availableCoins[i])
		}

		amountLeft = amountLeft % availableCoins[i]
	}

	return result
}
