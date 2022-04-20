package main

import "testing"

func TestSearchOnMap1(t *testing.T) {

	t.Run(">>> test if a solution for map1 can be found", func(t *testing.T) {
		var map1 [][]byte = [][]byte{
			{'+', '0', '+', '0', '+', '0', '0'},
			{'+', '+', '+', '0', '+', '0', '0'},
			{'0', '0', '0', '0', '0', '0', '0'},
			{'+', '0', '+', '0', '0', '+', '0'},
			{'+', '0', '0', '0', '+', '0', '+'},
			{'+', '+', '+', '+', '0', '0', '0'},
		}

		want := &position{row: 1, column: 4}
		got, err := searchNearestExit(map1, &position{row: 3, column: 1})
		if err != nil {
			t.Errorf("Unexpected error in the search: %s", err)
		}
		if want.row != got.row || want.column != got.column {
			t.Errorf("Error in search result: got (%d, %d), want (%d, %d)", got.row, got.column, want.row, want.column)
		}
	})
}
