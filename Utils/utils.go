package Utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func ReadFileByLine(input string) []string {

	f, err := os.Open(input)

	CheckError(err)

	defer f.Close()

	scanner := bufio.NewScanner(f)

	lines := make([]string, 0)

	for scanner.Scan() {

		lines = append(lines, scanner.Text())
	}

	CheckError(scanner.Err())

	return lines
}

func CheckError(err error) {

	if err != nil {
		log.Fatal(err)
	}

}

func StringSliceToIntSlice(slice []string) []int {

	intSlice := make([]int, len(slice))

	for i, s := range slice {

		if val, err := strconv.Atoi(s); err == nil {
			intSlice[i] = val
		}
	}

	return intSlice
}
