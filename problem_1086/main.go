////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Apr-19-2022  -  aldebap
//
//	Problem Sodoku Solver
////////////////////////////////////////////////////////////////////////////////

package main

type sodoku [9][9]uint

//	return true if the Sodoku matrix is a valid
func isSodokuSolved(input sodoku) bool {

	//	check Sodoku lines
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

	//	check Sodoku columns
	for i := 0; i < 9; i++ {
		var digit [9]bool

		for j := 0; j < 9; j++ {

			if input[j][i] == 0 {
				return false
			}
			if digit[input[j][i]-1] {
				return false
			}

			digit[input[j][i]-1] = true
		}
	}

	//	TODO: do I need to check each of 3x3 mini groups ?

	return true
}
