package main

import (
	"AOC_2021/Utils"
	"fmt"
)

func main() {

	lines := Utils.ReadFileByLine("./Day1/input.txt")

	//Part1
	count := checkIncreases(lines)
	fmt.Println(count)

	//Part2
	count = checkWindowIncreases(lines)
	fmt.Println(count)

}
