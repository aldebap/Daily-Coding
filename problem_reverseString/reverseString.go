////////////////////////////////////////////////////////////////////////////////
//	reverseString.go  -  Ago-19-2022  -  aldebap
//
//	Problem Reverse string
////////////////////////////////////////////////////////////////////////////////

package main

import "regexp"

//	reverseString return the reverse of a string (keeping symbold where they are)
func reverseString(input string) string {
	var result string

	//	regexp to match letters
	letter, err := regexp.Compile("^[a-zA-Z]$")
	if err != nil {
		panic("fail to compile regexp: " + err.Error())
	}

	//	find the last letter on input
	//	TODO: better to do this backwards !
	lastLetter := -1

	for i, char := range input {
		if letter.FindString(string(char)) != "" {
			lastLetter = i
		}
	}
	//	reverse the string
	for _, char := range input {
		if letter.FindString(string(char)) != "" {
			result = result + string(input[lastLetter])

			//	mode the lastLetter to the previous letter
			for {
				if lastLetter == 0 {
					break
				}

				lastLetter--
				if letter.FindString(string(input[lastLetter])) != "" {
					break
				}
			}
		} else {
			result = result + string(char)
		}
	}

	return result
}
