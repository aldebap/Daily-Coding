////////////////////////////////////////////////////////////////////////////////
//	main_test.go  -  Sep-1-2022  -  aldebap
//
//	Test cases for the Numerical Parser
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"fmt"
	"testing"
)

func TestIsNumeric(t *testing.T) {

	t.Run(">>> Test IsNumeric() parser", func(t *testing.T) {

		//	a few test cases
		var testScenarios = []struct {
			scenario string
			input    string
			output   bool
		}{
			{scenario: "Valid integer", input: "10", output: true},
			{scenario: "Valid positive integer", input: "+10", output: true},
			{scenario: "Valid negative integer", input: "-10", output: true},

			{scenario: "Invalid double signed positive integer", input: "++10", output: false},
			{scenario: "Invalid double signed negative integer", input: "--10", output: false},

			{scenario: "Valid real number", input: "10.1", output: true},
			{scenario: "Valid positive real number", input: "+10.1", output: true},
			{scenario: "Valid negative real number", input: "-10.1", output: true},

			{scenario: "Invalid real number without integer part", input: ".1", output: false},
			{scenario: "Invalid positive real number without integer part", input: "+.1", output: false},
			{scenario: "Invalid negative real number without integer part", input: "-.1", output: false},
			{scenario: "Invalid real number without decimal part", input: "10.", output: false},
			{scenario: "Invalid positive real number without decimal part", input: "+10.", output: false},
			{scenario: "Invalid negative real number without decimal part", input: "-10.", output: false},
			{scenario: "Invalid double signed positive real number", input: "++10.1", output: false},
			{scenario: "Invalid double signed negative real number", input: "--10.1", output: false},

			{scenario: "Valid cientific integer", input: "1e5", output: true},
			{scenario: "Valid positive cientific integer", input: "+1e5", output: true},
			{scenario: "Valid negative cientific integer", input: "-1e5", output: true},

			{scenario: "Valid cientific real number", input: "10.1e5", output: true},
			{scenario: "Valid positive cientific real number", input: "+10.1e5", output: true},
			{scenario: "Valid negative cientific real number", input: "-10.1e5", output: true},

			{scenario: "Invalid cientific integer without radix part", input: "e5", output: false},
			{scenario: "Invalid positive cientific integer without radix part", input: "+e5", output: false},
			{scenario: "Invalid negative cientific integer without radix part", input: "-e5", output: false},
			{scenario: "Invalid cientific integer without exponent part", input: "1e", output: false},
			{scenario: "Invalid positive cientific integer without exponent part", input: "+1e", output: false},
			{scenario: "Invalid negative cientific integer without exponent part", input: "-1e", output: false},
			{scenario: "Invalid double signed positive cientific integer", input: "++1e5", output: false},
			{scenario: "Invalid double signed negative cientific integer", input: "--1e5", output: false},

			{scenario: "Invalid character", input: "a", output: false},
			{scenario: "Invalid character before digits", input: "a1", output: false},
			{scenario: "Invalid character in the middle of digits", input: "1a0", output: false},
			{scenario: "Invalid character after digits", input: "10a", output: false},
			{scenario: "Invalid positive signal without digits", input: "+", output: false},
			{scenario: "Invalid negative signal without digits", input: "-", output: false},
		}

		for _, test := range testScenarios {

			fmt.Printf("scenario: %s\n", test.scenario)

			want := test.output
			got := IsNumeric(test.input)

			//	check the result
			if want != got {
				t.Errorf("failed parsing numeric string '%s': expected: %t result: %t", test.input, want, got)
			}
		}
	})
}
