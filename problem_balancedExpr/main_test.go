////////////////////////////////////////////////////////////////////////////////
//	main_test.go  -  Jan-31-2022  -  aldebap
//
//	Test cases for the Balanced Expression Problem
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
//	Balanced Expression test cases
////////////////////////////////////////////////////////////////////////////////

func TestExpressions(t *testing.T) {

	t.Run(">>> Balanced Expression Problem", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			input  string
			result bool
		}{
			{input: "", result: true},
			{input: "()", result: true},
			{input: "(", result: false},
			{input: ")", result: false},
			{input: "[]", result: true},
			{input: "[", result: false},
			{input: "]", result: false},
			{input: "{}", result: true},
			{input: "{", result: false},
			{input: "}", result: false},
			{input: "(()", result: false},
			{input: "())", result: false},
			{input: "([)", result: false},
			{input: "(])", result: false},
			{input: "({)", result: false},
			{input: "(})", result: false},
			{input: "([{}])", result: true},
			{input: "()[]{}", result: true},
			{input: "([{}])[{}]", result: true},
			{input: "([{}])[(}]", result: false},
		}

		for _, test := range testScenarios {

			want := test.result
			got := isBalanced(test.input)

			fmt.Printf("input: %s ; result: %t\n", test.input, got)
			if want != got {
				t.Errorf("input: %s ; expected: %t result: %t", test.input, want, got)
			}
		}
	})
}
