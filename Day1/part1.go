package main

import (
	"AOC_2021/Utils"
	"math"
)

func checkIncreases(lines []string) int {

	slice := Utils.StringSliceToIntSlice(lines)

	prevDepth := math.MaxInt
	count := 0

	for _, value := range slice {

		if value > prevDepth {
			count++
		}

		prevDepth = value

	}

	return count
}



