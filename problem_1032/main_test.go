////////////////////////////////////////////////////////////////////////////////
//	main_test.go  -  Feb-06-2022  -  aldebap
//
//	Problem Toeplitz matrix verification
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
//	Toeplitz matrix verification test cases
////////////////////////////////////////////////////////////////////////////////

func TestMatrix(t *testing.T) {

	t.Run(">>> Toeplitz matrix verification", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			matrix         [][]int32
			lines, columns uint
			result         bool
		}{
			{matrix: [][]int32{{1, 2, 3, 4, 8}, {5, 1, 2, 3, 4}, {4, 5, 1, 2, 3}, {7, 4, 5, 1, 2}}, lines: 4, columns: 5, result: true},
		}

		for _, test := range testScenarios {

			fmt.Printf(">>> [test] matrix: %v ; expected: %t", test.matrix, test.result)

			//	run matrix verification
			want := test.result
			got := isToeplitzMatrix(test.matrix, test.lines, test.columns)

			if want != got {
				t.Errorf("matrix: %v ; expected: %t result: %t", test.matrix, want, got)
			}
		}
	})
}
