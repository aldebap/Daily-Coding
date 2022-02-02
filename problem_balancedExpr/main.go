////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Jan-30-2022  -  aldebap
//
//	Problem Balanced Expression
////////////////////////////////////////////////////////////////////////////////

package main

//	return true if the input expression is balanced (non recursive implementation)
func isBalancedNonRecursive(expression string) bool {

	//fmt.Printf("[debug] checking expression: %s\n", expression)
	//	check for rule #1
	if len(expression) == 0 {
		return true
	}

	//	use a stack to make sure all closing delimiters are in the same order as they were open
	var delimiterStack []byte

	for i := 0; i < len(expression); i++ {
		if expression[i] == '(' || expression[i] == '[' || expression[i] == '{' {
			delimiterStack = append(delimiterStack, expression[i])
			continue
		}

		//	closing an unopened delimiter means expression is not balanced
		if (expression[i] == ')' || expression[i] == ']' || expression[i] == '}') && len(delimiterStack) == 0 {
			return false
		}

		//	check if a closing delimiter correspond to the previous opended
		if delimiterStack[len(delimiterStack)-1] == '(' && expression[i] == ')' {
			delimiterStack = delimiterStack[:len(delimiterStack)-1]
			continue
		}

		if delimiterStack[len(delimiterStack)-1] == '[' && expression[i] == ']' {
			delimiterStack = delimiterStack[:len(delimiterStack)-1]
			continue
		}

		if delimiterStack[len(delimiterStack)-1] == '{' && expression[i] == '}' {
			delimiterStack = delimiterStack[:len(delimiterStack)-1]
			continue
		}

		//	if we got here, we got an invalid character
		return false
	}

	if len(delimiterStack) == 0 {
		return true
	}

	return false
}

//	return true if the input expression is balanced (refactored implementation)
func isBalancedRefactored(expression string) bool {

	//fmt.Printf("[debug] checking expression: %s\n", expression)
	//	check for rule #1
	if len(expression) == 0 {
		return true
	}

	//	the valid expression delimiters
	var expressionDelimiter = []struct {
		openToken, closeToken byte
	}{
		{openToken: '(', closeToken: ')'},
		{openToken: '[', closeToken: ']'},
		{openToken: '{', closeToken: '}'},
	}

	for _, delimiter := range expressionDelimiter {

		if expression[0] == delimiter.openToken {

			openDelimiter := 1
			i := 1

			for ; i < len(expression); i++ {

				if expression[i] == delimiter.openToken {
					openDelimiter++
				}
				if expression[i] == delimiter.closeToken {
					openDelimiter--
				}
				//	when all opened delimiter of the same kind are closed, check the subexpressions: (A)B
				if openDelimiter == 0 {
					var subExprA, subExprB string

					if i > 1 && i < len(expression) {
						subExprA = expression[1:i]
					}
					if i+1 < len(expression) {
						subExprB = expression[i+1:]
					}

					//	check for rules #2 and #5 (avoiding too much recurrence when possible)
					return (len(subExprA) == 0 || isBalancedRefactored(subExprA)) && (len(subExprB) == 0 || isBalancedRefactored(subExprB))
				}
			}
		}
	}

	return false
}

//	return true if the input expression is balanced (the simplest implementation)
func isBalancedSimple(expression string) bool {

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

				return isBalancedSimple(subExprA) && isBalancedSimple(subExprB)
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

				return isBalancedSimple(subExprA) && isBalancedSimple(subExprB)
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

				return isBalancedSimple(subExprA) && isBalancedSimple(subExprB)
			}
		}
	}

	return false
}
