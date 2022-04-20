////////////////////////////////////////////////////////////////////////////////
//	main_test.go  -  Apr-19-2022  -  aldebap
//
//	Problem Sodoku Solver
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

////////////////////////////////////////////////////////////////////////////////
//	Sodoku Solver test cases
////////////////////////////////////////////////////////////////////////////////

func TestSodokuSolver(t *testing.T) {

	t.Run(">>> Sodoku Solver", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			sodokuGame sodoku
			result     bool
		}{
			{sodokuGame: sodoku{{3, 0, 2, 0, 8, 0, 1, 0, 5},
				{0, 0, 7, 2, 0, 0, 0, 6, 0},
				{5, 0, 8, 9, 0, 0, 0, 4, 7},
				{0, 8, 0, 4, 0, 0, 3, 0, 2},
				{0, 3, 0, 1, 6, 0, 0, 5, 8},
				{0, 1, 0, 5, 0, 0, 6, 7, 0},
				{0, 2, 0, 3, 4, 5, 0, 0, 1},
				{0, 0, 0, 0, 2, 6, 0, 0, 9},
				{0, 0, 0, 0, 9, 0, 5, 2, 6},
			}, result: true},
		}

		for _, test := range testScenarios {

			fmt.Printf(">>> [test] Sodoku: \n")
			printSodoku(test.sodokuGame)

			//	run matrix verification
			want := test.result
			got := isSodokuSolved(test.sodokuGame)

			if want != got {
				t.Errorf("expected: %t result: %t", want, got)
			}
		}
	})
}

func printSodoku(input sodoku) {

	fmt.Printf("+-------+-------+-------+\n")

	for i := 0; i < 9; i++ {
		fmt.Printf("| ")
		for j := 0; j < 9; j++ {
			if input[i][j] == 0 {
				fmt.Printf("  ")
			} else {
				fmt.Printf("%d ", input[i][j])
			}

			if (j+1)%3 == 0 {
				fmt.Printf("| ")
			}
		}
		fmt.Printf("\n")

		if (i+1)%3 == 0 {
			fmt.Printf("+-------+-------+-------+\n")
		}
	}
}
