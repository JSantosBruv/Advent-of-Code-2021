package main

import (
	"AOC_2021/Utils"
)

func checkWindowIncreases(lines []string) int {

	slice := Utils.StringSliceToIntSlice(lines)
	prevCount := make([]int, 3)
	count := 0

	//Initialize counts with the three first positions
	for i, val := range slice[:3] {

		prevCount[0] += val

		if i >= 1 {
			prevCount[1] += val
		}
		if i == 2 {
			prevCount[2] = val
		}
	}

	for _, val := range slice[3:] {

		prevCount[1] += val
		prevCount[2] += val

		if prevCount[1] > prevCount[0] {
			count++
		}

		prevCount[0] = prevCount[1]
		prevCount[1] = prevCount[2]
		prevCount[2] = val
	}

	return count
}
