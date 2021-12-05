package main

import (
	"strconv"
	"strings"
)

func submarine(lines []string) int {

	horizontal, depth := 0, 0

	for _, move := range lines {

		info := strings.Split(move, " ")

		val, _ := strconv.Atoi(info[1])

		switch info[0] {

		case forward:
			horizontal += val
		case down:
			depth += val
		case up:
			depth -= val

		}
	}

	return horizontal * depth
}
