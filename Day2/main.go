package main

import (
	"AOC_2021/Utils"
	"fmt"
)

const (
	forward = "forward"
	up      = "up"
	down    = "down"
)

func main() {

	lines := Utils.ReadFileByLine("./Day2/input.txt")

	//Part1
	result := submarine(lines)
	fmt.Println(result)

	//Part2
	result = submarineAim(lines)
	fmt.Println(result)

}
