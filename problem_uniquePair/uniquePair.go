////////////////////////////////////////////////////////////////////////////////
//	uniquePair.go  -  Sep-8-2022  -  aldebap
//
//	Problem Unique Pair
////////////////////////////////////////////////////////////////////////////////

package main

import (
	"math"
)

//	uniquePair return a slice with all unique pairs in input array
func uniquePair(input []int) []int {

	//	search for the unique pairs
	pair := make(map[int]bool)

	for _, num := range input {
		_, exists := pair[num]
		if !exists {
			_, exists := pair[-num]
			if !exists {
				pair[num] = false
			} else {
				pair[-num] = true
			}
		}
	}

	//	all numbers that are mapped to true are unique pairs so, append them to the result slice
	result := make([]int, 0)

	for num := range pair {
		if pair[num] {
			result = append(result, int(math.Abs(float64(num))))
		}
	}

	return result
}
