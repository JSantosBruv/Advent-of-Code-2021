package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func readFileByLine() []string {

	f, err := os.Open("./Day1/input.txt")

	checkError(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)

	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	checkError(scanner.Err())

	return lines
}

func checkIncreases(lines []string) int {

	prevDepth := math.MaxInt
	count := 0

	for _, line := range lines {

		currDepth, err := strconv.Atoi(line)

		checkError(err)

		if currDepth > prevDepth {
			count++
		}

		prevDepth = currDepth

	}

	return count
}

func main() {

	lines := readFileByLine()

	count := checkIncreases(lines)

	fmt.Print(count)

}

func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}

}
