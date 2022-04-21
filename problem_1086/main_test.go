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

		type testData struct {
			sodokuGame sodoku
			result     bool
		}
		//	a few test cases
		var testScenarios []testData = []testData{
			//	two easy level sodokus
			{sodokuGame: sodoku{
				{3, 0, 2, 0, 8, 0, 1, 0, 5},
				{0, 0, 7, 2, 0, 0, 0, 6, 0},
				{5, 0, 8, 9, 0, 0, 0, 4, 7},
				{0, 8, 0, 4, 0, 0, 3, 0, 2},
				{0, 3, 0, 1, 6, 0, 0, 5, 8},
				{0, 1, 0, 5, 0, 0, 6, 7, 0},
				{0, 2, 0, 3, 4, 5, 0, 0, 1},
				{0, 0, 0, 0, 2, 6, 0, 0, 9},
				{0, 0, 0, 0, 9, 0, 5, 2, 6},
			}, result: true,
			},
			{sodokuGame: sodoku{
				{0, 0, 4, 0, 0, 7, 8, 3, 0},
				{0, 1, 0, 9, 3, 0, 7, 0, 5},
				{3, 5, 7, 0, 0, 4, 1, 0, 0},
				{0, 3, 0, 0, 0, 2, 9, 8, 4},
				{0, 0, 0, 0, 8, 1, 0, 0, 0},
				{0, 2, 0, 0, 0, 0, 6, 0, 0},
				{4, 7, 3, 2, 0, 8, 0, 0, 1},
				{2, 6, 0, 1, 0, 0, 0, 7, 8},
				{5, 0, 1, 6, 0, 0, 4, 0, 0},
			}, result: true,
			},
		}

		for _, test := range testScenarios {

			fmt.Printf(">>> [test] Sodoku: \n")
			print(&test.sodokuGame)

			//	try to solve the Sodoku
			solution, err := solve(&test.sodokuGame)
			if err != nil {
				t.Errorf("error: %s", err)
			} else {
				fmt.Printf("[test] solution: \n")
				print(solution)
			}

			want := test.result
			got := isSolved(solution)
			if want != got {
				t.Errorf("expected: %t result: %t", want, got)
			}
		}
	})
}
