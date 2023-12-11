package _9

import (
	"advent_of_code23/utils"
	"regexp"
	"strconv"
	"strings"
)

func Problem1(inputFileName string, rev bool) int {
	lines := strings.Split(utils.GetFileAsString(inputFileName), "\n")

	sum := 0
	for l := range lines {
		line := getLineAsInts(lines[l])
		sum += getNextR(line, rev)
	}

	return sum
}

func getNextR(v []int, rev bool) int {
	vals := make([]int, len(v)-1)
	hasNonZero := false
	for i := 1; i < len(v); i++ {
		difference := v[i] - v[i-1]
		vals[i-1] = difference
		if difference != 0 {
			hasNonZero = true
		}
	}
	if !rev {
		if hasNonZero {
			return v[len(v)-1] + getNextR(vals, rev)
		} else {
			return v[len(v)-1]
		}
	} else {
		if hasNonZero {
			return v[0] - getNextR(vals, rev)
		} else {
			return v[0]
		}

	}

}

func getLineAsInts(line string) []int {
	val := regexp.MustCompile(`[\d-]+`).FindAllString(line, -1)

	returnVal := make([]int, len(val))
	for i := range val {
		v, err := strconv.Atoi(val[i])
		if err != nil {
			panic(err)
		}
		returnVal[i] = v
	}

	return returnVal
}
