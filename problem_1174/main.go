////////////////////////////////////////////////////////////////////////////////
//	main.go  -  Sep-1-2022  -  aldebap
//
//	Problem Numerical Parser
////////////////////////////////////////////////////////////////////////////////

package main

import "regexp"

var (
	integerRegex          *regexp.Regexp
	floatRegex            *regexp.Regexp
	cientificIntegerRegex *regexp.Regexp
	cientificFloatRegex   *regexp.Regexp
)

//	init initialize the regexs uses for the parser
func init() {

	var err error

	//	regexp to match integers
	integerRegex, err = regexp.Compile("^([+-]{0,1}[0-9]+)$")
	if err != nil {
		panic("fail to compile regexp: " + err.Error())
	}

	//	regexp to match floats
	floatRegex, err = regexp.Compile("^([+-]{0,1}[0-9]+[.][0-9]+)$")
	if err != nil {
		panic("fail to compile regexp: " + err.Error())
	}

	//	regexp to match cientific integers
	cientificIntegerRegex, err = regexp.Compile("^([+-]{0,1}[0-9]+[eE][0-9]+)$")
	if err != nil {
		panic("fail to compile regexp: " + err.Error())
	}

	//	regexp to match cientific floats
	cientificFloatRegex, err = regexp.Compile("^([+-]{0,1}[0-9]+[.][0-9]+[eE][0-9]+)$")
	if err != nil {
		panic("fail to compile regexp: " + err.Error())
	}
}

//	IsNumeric returns true if the input represents a numerical value
func IsNumeric(input string) bool {

	//	is an integer ?
	if integerRegex.FindString(input) != "" {
		return true
	}

	//	is a float ?
	if floatRegex.FindString(input) != "" {
		return true
	}

	//	is an cientific integer ?
	if cientificIntegerRegex.FindString(input) != "" {
		return true
	}

	//	is a cientific float ?
	if cientificFloatRegex.FindString(input) != "" {
		return true
	}

	return false
}
