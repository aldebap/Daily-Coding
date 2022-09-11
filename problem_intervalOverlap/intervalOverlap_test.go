////////////////////////////////////////////////////////////////////////////////
//	intervalOverlap.go  -  Sep-9-2022  -  aldebap
//
//	Test cases for the Date invervals overlap
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"testing"
)

//	TestCheckOverlap check overlap test cases
func TestCheckOverlap(t *testing.T) {

	t.Run(">>> Reverse Unique Pair", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			existing []dateInterval
			incoming dateInterval
			result   bool
		}{
			{existing: []dateInterval{}, incoming: dateInterval{start: "2020-01-06", end: "2020-01-09"}, result: false},
			{existing: []dateInterval{{start: "2020-01-01", end: "2020-01-05"}}, incoming: dateInterval{start: "2020-01-06", end: "2020-01-09"}, result: false},
			{existing: []dateInterval{{start: "2020-01-10", end: "2020-01-15"}}, incoming: dateInterval{start: "2020-01-06", end: "2020-01-09"}, result: false},

			{existing: []dateInterval{{start: "2020-01-01", end: "2020-01-07"}}, incoming: dateInterval{start: "2020-01-06", end: "2020-01-09"}, result: true},
			{existing: []dateInterval{{start: "2020-01-08", end: "2020-01-10"}}, incoming: dateInterval{start: "2020-01-06", end: "2020-01-09"}, result: true},
			{existing: []dateInterval{{start: "2020-01-07", end: "2020-01-08"}}, incoming: dateInterval{start: "2020-01-06", end: "2020-01-09"}, result: true},

			{existing: []dateInterval{
				{start: "2020-01-01", end: "2020-01-05"},
				{start: "2020-01-10", end: "2020-01-20"},
				{start: "2020-01-07", end: "2020-01-08"}},
				incoming: dateInterval{start: "2020-01-06", end: "2020-01-09"}, result: true},
		}

		for _, test := range testScenarios {

			//	test the implementation
			want := test.result
			got, err := CheckOverlap(test.incoming, test.existing)

			if err != nil {
				t.Errorf("failed checking for overlaping: %s", err.Error())
			}

			if want != got {
				t.Errorf("test scenario failed: expected result %t, got %t", want, got)
			}
		}
	})
}
