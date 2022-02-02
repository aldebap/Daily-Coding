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

			fmt.Printf("input: '%s'\texpected: %t\n", test.input, test.result)

			//	test the simple implementation
			want := test.result
			got := isBalancedSimple(test.input)

			if want != got {
				t.Errorf("input: %s ; expected: %t result: %t", test.input, want, got)
			}

			//	test the refactored implementation
			want = test.result
			got = isBalancedRefactored(test.input)

			if want != got {
				t.Errorf("input: %s ; expected: %t result: %t", test.input, want, got)
			}

			//	test the non recursive implementation
			want = test.result
			got = isBalancedNonRecursive(test.input)

			if want != got {
				t.Errorf("input: %s ; expected: %t result: %t", test.input, want, got)
			}
		}
	})
}
