////////////////////////////////////////////////////////////////////////////////
//	wordPuzzle_test.go  -  Sep-12-2022  -  aldebap
//
//	Test cases for the word puzzle games
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"testing"
)

//	TestWordPuzzleGame test cases to validate all searches made in existsInWordPuzzle()
func TestWordPuzzleGame(t *testing.T) {

	t.Run(">>> word puzzle game", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			wordPuzzle [][]byte
			search     string
			result     bool
		}{
			{wordPuzzle: [][]byte{{'A', 'U', 'I', 'K', 'F', 'W', 'N'},
				{'W', 'Q', 'B', 'O', 'L', 'D', 'P'},
				{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
				{'Y', 'Z', 'X', 'E', 'R', 'L', 'W'}},
				search: "BOLD", result: true},
			{wordPuzzle: [][]byte{{'A', 'U', 'I', 'K', 'F', 'W', 'N'},
				{'W', 'Q', 'B', 'O', 'L', 'X', 'P'},
				{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
				{'Y', 'Z', 'X', 'E', 'R', 'L', 'W'}},
				search: "REAL", result: true},

			{wordPuzzle: [][]byte{{'A', 'U', 'I', 'K', 'F', 'W', 'N'},
				{'W', 'Q', 'B', 'O', 'L', 'X', 'P'},
				{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
				{'Y', 'Z', 'R', 'E', 'R', 'L', 'W'}},
				search: "BAR", result: true},
			{wordPuzzle: [][]byte{{'A', 'U', 'I', 'K', 'F', 'T', 'N'},
				{'W', 'Q', 'B', 'O', 'L', 'F', 'P'},
				{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
				{'Y', 'Z', 'X', 'E', 'R', 'L', 'W'}},
				search: "LEFT", result: true},

			{wordPuzzle: [][]byte{{'A', 'U', 'I', 'K', 'F', 'W', 'N'},
				{'W', 'Q', 'B', 'O', 'L', 'X', 'P'},
				{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
				{'Y', 'Z', 'X', 'E', 'R', 'L', 'W'}},
				search: "UBER", result: true},
			{wordPuzzle: [][]byte{{'A', 'U', 'I', 'D', 'F', 'W', 'N'},
				{'W', 'Q', 'B', 'O', 'L', 'X', 'P'},
				{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
				{'Y', 'Z', 'X', 'E', 'R', 'L', 'W'}},
				search: "WELD", result: true},

			{wordPuzzle: [][]byte{{'A', 'U', 'I', 'K', 'F', 'W', 'N'},
				{'W', 'Q', 'B', 'T', 'L', 'X', 'P'},
				{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
				{'Y', 'C', 'X', 'E', 'R', 'L', 'W'}},
				search: "CAT", result: true},
			{wordPuzzle: [][]byte{{'A', 'U', 'I', 'K', 'F', 'W', 'N'},
				{'W', 'Q', 'B', 'O', 'L', 'X', 'P'},
				{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
				{'Y', 'M', 'X', 'E', 'R', 'L', 'W'}},
				search: "FOAM", result: true},

			{wordPuzzle: [][]byte{{'A', 'U', 'I', 'K', 'F', 'W', 'N'},
				{'W', 'Q', 'B', 'O', 'L', 'X', 'P'},
				{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
				{'Y', 'Z', 'X', 'E', 'R', 'L', 'W'}},
				search: "PET", result: false},
			{wordPuzzle: [][]byte{{'A', 'U', 'I', 'K', 'F', 'W', 'N'},
				{'W', 'Q', 'B', 'O', 'L', 'X', 'P'},
				{'T', 'L', 'A', 'E', 'R', 'E', 'S'},
				{'Y', 'Z', 'X', 'E', 'R', 'L', 'W'}},
				search: "HUMANITY", result: false},
		}

		for _, test := range testScenarios {

			//	test the implementation
			want := test.result
			got := ExistsInWordPuzzle(test.search, test.wordPuzzle)

			if want != got {
				t.Errorf("test scenario failed: expected result %t for word %s, got %t", want, test.search, got)
			}
		}
	})
}
