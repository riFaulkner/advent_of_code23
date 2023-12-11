package _8

import (
	"advent_of_code23/utils"
	"regexp"
	"slices"
	"strings"
)

type Instruction struct {
	left  string
	right string
}

func Problem1(inputFileName string) int {
	segments := strings.Split(utils.GetFileAsString(inputFileName), "\n\n")

	maps := generateMap(&segments[1])

	return findLocation("AAA", "ZZZ", segments[0], maps)
}

func Problem2(inputFileName string) int {
	segments := strings.Split(utils.GetFileAsString(inputFileName), "\n\n")

	maps := generateMap(&segments[1])
	locations := getAllStartingLocations('A', &maps)
	l := make([]int, len(locations))

	for i := range locations {
		l[i] = findLocation(locations[i], "Z", segments[0], maps)
	}

	return findLowestCommonMultiple(l)
}

func findLocation(start, end string, directions string, maps map[string]Instruction) int {
	i := 0
	loc := start

	directionsLength := len(directions)
	for {
		if len(end) == 1 {
			if loc[len(loc)-1] == end[0] {
				return i
			}
		} else {
			if loc == end {
				return i
			}
		}

		if directions[i%directionsLength] == 'R' {
			i++
			loc = maps[loc].right
		} else {
			i++
			loc = maps[loc].left
		}
	}
}

func findLowestCommonMultiple(vals []int) int {
	runningVals := slices.Clone(vals)
	high := 0
	for {
		matches := 0

		for j, _ := range runningVals {
			if runningVals[j] > high {
				high = runningVals[j]
				matches = 1
				continue
			}
			if runningVals[j] == high {
				matches++
				continue
			}
			runningVals[j] += vals[j]
		}

		if matches == len(runningVals) {
			return high
		}
	}
}
func getAllStartingLocations(key byte, maps *map[string]Instruction) []string {
	var result []string
	for k, _ := range *maps {
		if k[len(k)-1] == key {
			result = append(result, k)
		}
	}
	return result
}

func generateMap(input *string) map[string]Instruction {
	result := make(map[string]Instruction)
	for _, line := range strings.Split(*input, "\n") {
		ins := regexp.MustCompile(`[A-Z\d]+`).FindAllString(line, -1)
		result[ins[0]] = Instruction{left: ins[1], right: ins[2]}
	}
	return result
}
