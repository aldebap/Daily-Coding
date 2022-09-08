////////////////////////////////////////////////////////////////////////////////
//	uniquePair_test.go  -  Sep-8-2022  -  aldebap
//
//	Test cases for the Unique Pair Problem
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

//	TestExpressions Reverse string test cases
func TestUniquePair(t *testing.T) {

	t.Run(">>> Reverse Unique Pair", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			input  []int
			output []int
		}{
			{input: []int{}, output: []int{}},
			{input: []int{1, 2, 3, 4, 5}, output: []int{}},
			{input: []int{1, 2, 3, 4, 5, -1, -3, -5}, output: []int{1, 3, 5}},
			{input: []int{-1, -2, -3, -4, -5, 1, 3, 5}, output: []int{1, 3, 5}},
			{input: []int{7, -5, 6, 5, -8, 5, -5, 1, 7, 4, -1, -2, 1}, output: []int{5, 1}},
		}

		for _, test := range testScenarios {

			fmt.Printf("input: %v\texpected: %v\n", test.input, test.output)

			//	test the implementation
			want := test.output
			got := uniquePair(test.input)

			if len(want) != len(got) {
				t.Errorf("failed finding unique pair: %v - expected result set size: %d result: %d", test.input, len(want), len(got))
			}

			for _, outputNum := range test.output {
				found := false

				for _, resultNum := range got {
					if outputNum == resultNum {
						found = true
						break
					}
				}

				if !found {
					t.Errorf("failed finding unique pair: %d expected in the result set: %v", outputNum, got)
				}
			}
		}
	})
}
