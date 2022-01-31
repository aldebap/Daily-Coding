////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Jan-30-2022  -  aldebap
//
//	Problem Balanced Expression
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
	fmt.Printf(">>> Problem Balanced Expression\n\n")

	//	a few test cases
	type testCase struct {
		input  string
		result bool
	}

	var testScenarios []testCase = []testCase{
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
		{input: "([)", result: false},
		{input: "({)", result: false},
		{input: "([{}])", result: true},
		{input: "()[]{}", result: true},
		{input: "([{}])[{}]", result: true},
		{input: "([{}])[(}]", result: false},
	}

	for _, test := range testScenarios {

		fmt.Printf("input: %s ; expected: %t; result: %t\n", test.input, test.result, isBalanced(test.input))
	}
}

//	return true if the input expression is balanced
func isBalanced(expression string) bool {

	//fmt.Printf("[debug] checking expression: %s\n", expression)
	if len(expression) == 0 {
		return true
	}

	//	subexpression between parenthesis
	if expression[0] == '(' {

		openParentesis := 1
		i := 1

		for ; i < len(expression); i++ {

			if expression[i] == '(' {
				openParentesis++
			}
			if expression[i] == ')' {
				openParentesis--
			}
			//	when all opened parenthesis are closed, check the subexpressions: (A)B
			if openParentesis == 0 {
				subExprA := ""
				subExprB := ""

				if i > 1 && i < len(expression) {
					subExprA = expression[1:i]
				}
				if i+1 < len(expression) {
					subExprB = expression[i+1:]
				}

				return isBalanced(subExprA) && isBalanced(subExprB)
			}
		}
	}

	//	subexpression between square brackets
	if expression[0] == '[' {

		openSquareBrackets := 1
		i := 1

		for ; i < len(expression); i++ {

			if expression[i] == '[' {
				openSquareBrackets++
			}
			if expression[i] == ']' {
				openSquareBrackets--
			}
			//	when all opened square brackets are closed, check the subexpressions: [A]B
			if openSquareBrackets == 0 {
				subExprA := ""
				subExprB := ""

				if i > 1 && i < len(expression) {
					subExprA = expression[1:i]
				}
				if i+1 < len(expression) {
					subExprB = expression[i+1:]
				}

				return isBalanced(subExprA) && isBalanced(subExprB)
			}
		}
	}

	//	subexpression between brackets
	if expression[0] == '{' {

		openBrackets := 1
		i := 1

		for ; i < len(expression); i++ {

			if expression[i] == '{' {
				openBrackets++
			}
			if expression[i] == '}' {
				openBrackets--
			}
			//	when all opened brackets are closed, check the subexpressions: {A}B
			if openBrackets == 0 {
				subExprA := ""
				subExprB := ""

				if i > 1 && i < len(expression) {
					subExprA = expression[1:i]
				}
				if i+1 < len(expression) {
					subExprB = expression[i+1:]
				}

				return isBalanced(subExprA) && isBalanced(subExprB)
			}
		}
	}

	return false
}
