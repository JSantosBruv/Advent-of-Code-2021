package main

import (
	"strconv"
	"strings"
)

func submarineAim(lines []string) int {

	horizontal, depth, aim := 0, 0, 0

	for _, move := range lines {

		info := strings.Split(move, " ")

		val, _ := strconv.Atoi(info[1])

		switch info[0] {

		case forward:
			depth += aim * val
			horizontal += val

		case down:
			aim += val

		case up:
			aim -= val

		}
	}

	return horizontal * depth
}
