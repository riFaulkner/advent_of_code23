package _8

import (
	"advent_of_code23/utils"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"time"
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

// TODO: bench mark this to find ways to make it faster
func Problem2(inputFileName string) int {
	start := time.Now()
	fmt.Printf("starting problem 2 \n")
	segments := strings.Split(utils.GetFileAsString(inputFileName), "\n\n")
	fmt.Printf("generating map %v \n", time.Since(start))

	maps := generateMap(&segments[1])
	fmt.Printf("finding locations %v \n", time.Since(start))
	locations := getAllStartingLocations('A', &maps)
	l := make([]int, len(locations))
	fmt.Printf("Starting to find locations %v \n", time.Since(start))
	for i := range locations {
		l[i] = findLocation(locations[i], "Z", segments[0], maps)
	}
	fmt.Printf("finding lowest common multiple %v \n", time.Since(start))
	//return findLowestCommonMultiple(l)
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
	high := slices.Max(vals)

	for i := 1; ; i++ {
		v := high * i
		for j := range vals {
			if v%vals[j] != 0 {
				break
			}
			if j == len(vals)-1 {
				return v
			}
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
