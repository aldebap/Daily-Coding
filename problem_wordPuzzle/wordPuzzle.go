////////////////////////////////////////////////////////////////////////////////
//	wordPuzzle.go  -  Sep-12-2022  -  aldebap
//
//	Problem of word puzzle games
////////////////////////////////////////////////////////////////////////////////

package main

import "fmt"

//	ExistsInWordPuzzle returns true if the search word exists in the word puzzle
func ExistsInWordPuzzle(word string, puzzle [][]byte) bool {

	fmt.Printf("[debug] word: %s\n", word)
	//  first, same line search
	for _, line := range puzzle {
		//  left to right
		for i, _ := range line {
			found := true
			for j, _ := range word {
				if i+j > len(line)-1 {
					found = false
					break
				}
				if line[i+j] != word[j] {
					found = false
					break
				}
			}
			if found {
				fmt.Printf("[debug] left to right\n")
				return true
			}
		}

		//  right to left
		for i := len(line) - 1; i >= 0; i-- {
			found := true
			for j, _ := range word {
				if i-j < 0 {
					found = false
					break
				}
				if line[i-j] != word[j] {
					found = false
					break
				}
			}
			if found {
				fmt.Printf("[debug] right to left\n")
				return true
			}
		}
	}

	//  second, same column search
	for i, _ := range puzzle[0] {
		//  top to bottom
		for j, _ := range puzzle {
			found := true
			for k, _ := range word {
				if j+k > len(puzzle)-1 {
					found = false
					break
				}
				if puzzle[j+k][i] != word[k] {
					found = false
					break
				}
			}
			if found {
				fmt.Printf("[debug] top to bottom\n")
				return true
			}
		}

		//  bottom to top
		for j := len(puzzle) - 1; j >= 0; j-- {
			found := true
			for k, _ := range word {
				if j-k < 0 {
					found = false
					break
				}
				if puzzle[j-k][i] != word[k] {
					found = false
					break
				}
			}
			if found {
				fmt.Printf("[debug] bottom to top\n")
				return true
			}
		}
	}

	//  third, back slash like diagonal
	for i, _ := range puzzle[0] {
		//  top-left to bottom-rigth
		for j, _ := range puzzle {
			found := true
			for k, _ := range word {
				if j+k > len(puzzle)-1 {
					found = false
					break
				}
				if i+k > len(puzzle[0])-1 {
					found = false
					break
				}
				if puzzle[j+k][i+k] != word[k] {
					found = false
					break
				}
			}
			if found {
				fmt.Printf("[debug] top-left to bottom-rigth\n")
				return true
			}
		}

		//  bottom-rigth to top-left
		for j := len(puzzle) - 1; j >= 0; j-- {
			found := true
			for k, _ := range word {
				if j-k < 0 {
					found = false
					break
				}
				if i-k < 0 {
					found = false
					break
				}
				if puzzle[j-k][i-k] != word[k] {
					found = false
					break
				}
			}
			if found {
				fmt.Printf("[debug] bottom-rigth to top-left\n")
				return true
			}
		}
	}

	//  fourth, slash like diagonal
	for i, _ := range puzzle[0] {
		//  bottom-left to top-rigth
		for j := len(puzzle) - 1; j >= 0; j-- {
			found := true
			for k, _ := range word {
				if j-k < 0 {
					found = false
					break
				}
				if i+k > len(puzzle[0])-1 {
					found = false
					break
				}
				if puzzle[j-k][i+k] != word[k] {
					found = false
					break
				}
			}
			if found {
				fmt.Printf("[debug] bottom-left to top-rigth\n")
				return true
			}
		}

		//  top-rigth to bottom-left
		for j, _ := range puzzle {
			found := true
			for k, _ := range word {
				if j+k > len(puzzle)-1 {
					found = false
					break
				}
				if i-k < 0 {
					found = false
					break
				}
				if puzzle[j+k][i-k] != word[k] {
					found = false
					break
				}
			}
			if found {
				fmt.Printf("[debug] top-rigth to bottom-left\n")
				return true
			}
		}
	}

	return false
}
