////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Feb-05-2022  -  aldebap
//
//	Problem Toeplitz matrix verification
////////////////////////////////////////////////////////////////////////////////

package main

//	return true if the input matrix is a Toeplitz
func isToeplitzMatrix(matrix [][]int32, lines, columns uint) bool {

	//	check above the matrix diagonal
	for i := uint(0); i < columns; i++ {
		for j := uint(1); j < lines; j++ {

			if i+j < columns {
			}
			if i+j < columns && matrix[0][i] != matrix[j][i+j] {
				return false
			}
		}
	}

	//	check below the matrix diagonal
	for i := uint(1); i < lines; i++ {
		for j := uint(1); j < columns; j++ {

			if i+j < lines {
			}
			if i+j < lines && matrix[i][0] != matrix[i+j][j] {
				return false
			}
		}
	}

	return true
}
