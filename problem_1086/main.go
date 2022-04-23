////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Apr-19-2022  -  aldebap
//
//	Problem Sodoku Solver
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"errors"
	"fmt"
)

type sodoku [9][9]uint

//	solve a input sodoku and return the solution
func solve(input *sodoku) (*sodoku, error) {

	var solution sodoku

	//	copy the input Sodoku to the solution
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			solution[i][j] = input[i][j]
		}
	}

	//	first try to solve it by logic

	for {
		var cellsSolved = 0

		if isSolved(&solution) {
			break
		}

		//	for each empty cell (zero), try to check if there's only one digit that can be assigned to it
		for i := 0; i < 9; i++ {
			for j := 0; j < 9; j++ {
				if solution[i][j] == 0 {
					var usedDigit [9]bool

					//	check in the same line for digits already used
					for k := 0; k < 9; k++ {
						if j != k && solution[i][k] != 0 {
							usedDigit[solution[i][k]-1] = true
						}
					}

					//	check in the same column for digits already used
					for k := 0; k < 9; k++ {
						if i != k && solution[k][j] != 0 {
							usedDigit[solution[k][j]-1] = true
						}
					}

					//	check in the same box for digits already used
					var boxLine = i / 3
					var boxColumn = j / 3

					for k := 0; k < 3; k++ {
						for l := 0; l < 3; l++ {
							if i != 3*boxLine+k && j != 3*boxColumn+l && solution[3*boxLine+k][3*boxColumn+l] != 0 {
								usedDigit[solution[3*boxLine+k][3*boxColumn+l]-1] = true
							}
						}
					}

					//	check the two boxes in the same line for digits already used

					//	check the two boxes in the same column for digits already used

					//	if only one digit is not used, cell is solved
					var usedDigits = 0

					for k := 0; k < 9; k++ {
						if usedDigit[k] {
							usedDigits++
						}
					}
					if usedDigits == 8 {
						for k := 0; k < 9; k++ {
							if !usedDigit[k] {
								fmt.Printf("[trace] cell (%d, %d) set to %d\n", i+1, j+1, k+1)
								solution[i][j] = uint(k + 1)
								cellsSolved++
							}
						}
					}
				}
			}
		}

		//	if couldn't solve any cell, we got into an error
		if cellsSolved == 0 {
			return nil, errors.New("cannot solve this Sodoku")
		}
	}

	return &solution, nil
}

//	return true if the Sodoku matrix is a valid
func isSolved(input *sodoku) bool {

	//	check if Sodoku lines have empty space or duplicate digits
	for i := 0; i < 9; i++ {
		var digit [9]bool

		for j := 0; j < 9; j++ {

			if input[i][j] == 0 {
				return false
			}
			if digit[input[i][j]-1] {
				return false
			}

			digit[input[i][j]-1] = true
		}
	}

	//	check if Sodoku columns have duplicate digits
	for i := 0; i < 9; i++ {
		var digit [9]bool

		for j := 0; j < 9; j++ {

			if digit[input[j][i]-1] {
				return false
			}

			digit[input[j][i]-1] = true
		}
	}

	//	check in the same box for digits already used
	for boxLine := 0; boxLine < 3; boxLine++ {
		for boxColumn := 0; boxColumn < 3; boxColumn++ {
			var digit [9]bool

			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if digit[input[3*boxLine+i][3*boxColumn+j]-1] {
						return false
					}

					digit[input[3*boxLine+i][3*boxColumn+j]-1] = true
				}
			}
		}
	}

	return true
}

//	print a sodoku to the stdout
func print(input *sodoku) {

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
