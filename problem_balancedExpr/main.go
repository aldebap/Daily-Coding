////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Jan-30-2022  -  aldebap
//
//	Problem Balanced Expression
////////////////////////////////////////////////////////////////////////////////

package main

//	return true if the input expression is balanced
func isBalanced(expression string) bool {

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
				//	when all opened parenthesis are closed, check the subexpressions: (A)B
				if openDelimiter == 0 {
					var subExprA, subExprB string

					if i > 1 && i < len(expression) {
						subExprA = expression[1:i]
					}
					if i+1 < len(expression) {
						subExprB = expression[i+1:]
					}

					//	check for rules #2 and #5
					return (len(subExprA) == 0 || isBalanced(subExprA)) && (len(subExprB) == 0 || isBalanced(subExprB))
				}
			}
		}
	}

	return false
}
