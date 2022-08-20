////////////////////////////////////////////////////////////////////////////////
//	reverseString_test.go  -  Ago-19-2022  -  aldebap
//
//	Test cases for the Reverse string Problem
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
//	Reverse string test cases
////////////////////////////////////////////////////////////////////////////////

func TestExpressions(t *testing.T) {

	t.Run(">>> Reverse string Problem", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			input  string
			output string
		}{
			{input: "", output: ""},
			{input: "a", output: "a"},
			{input: "$", output: "$"},
			{input: "ab-cd", output: "dc-ba"},
			{input: "shg?ikr!rufaw++", output: "waf?urr!kighs++"},
		}

		for _, test := range testScenarios {

			fmt.Printf("input: '%s'\texpected: %s\n", test.input, test.output)

			//	test the implementation
			want := test.output
			got := reverseString(test.input)

			if want != got {
				t.Errorf("failed reversing input: %s ; expected: %s result: %s", test.input, want, got)
			}
		}
	})
}
