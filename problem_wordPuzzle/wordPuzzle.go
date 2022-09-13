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
		//  top to down
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
				fmt.Printf("[debug] top to down\n")
				return true
			}
		}

		//  down to top
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
				fmt.Printf("[debug] down to top\n")
				return true
			}
		}
	}

	//  third, back slash like diagonal
	for i, _ := range puzzle[0] {
		//  top-left to down-rigth
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
				fmt.Printf("[debug] top-left to down-rigth\n")
				return true
			}
		}

		//  down-rigth to top-left
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
				fmt.Printf("[debug] down-rigth to top-left\n")
				return true
			}
		}
	}

	//  fourth, slash like diagonal

	return false
}
