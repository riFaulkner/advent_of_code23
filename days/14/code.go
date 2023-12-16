package _14

import (
	"advent_of_code23/utils"
	"bytes"
	"fmt"
	"slices"
)

func Problem1(inputFileName string) int {
	grid := bytes.Split(utils.GetFileContent(inputFileName), []byte("\n"))

	shiftGrid(&grid, 'N')

	return getGridWeight(&grid)
}

func Problem2(inputFileName string) int {
	grid := bytes.Split(utils.GetFileContent(inputFileName), []byte("\n"))

	cycleLength := 0
	history := make([]int, 0)

	for i := 0; i < 1_000_000_000; i++ {
		performCycle(&grid)
		weight := getGridWeight(&grid)
		history = append(history, weight)
		if c := checkForCycle(history); c != -1 {
			cycleLength = c
			fmt.Printf("Ran %d cycles to find a cycle of length %d\n", i, cycleLength)
			break
		}
	}

	remainder := len(history) - (3 * cycleLength)
	remainder = (1_000_000_000 - remainder) % cycleLength

	fmt.Printf("Remainder: %d\n", remainder)

	for i := 0; i < remainder; i++ {
		performCycle(&grid)
	}
	fmt.Printf("Final cycle weight: %d\n", getGridWeight(&grid))

	return getGridWeight(&grid)
}

func checkForCycle(history []int) int {
	numRepeats := 3
	historyLength := len(history)
	if historyLength < numRepeats*2 {
		return -1
	}
	for i := 1; i <= historyLength/numRepeats; i++ {
		s1 := history[historyLength-i:]                        // final i numbers in the slice
		s2 := history[historyLength-i-i : historyLength-i]     // the final i numbers before the final i numbers
		s3 := history[historyLength-i-i-i : historyLength-i-i] // the final i numbers before the final i numbers

		if slices.Equal(s1, s2) && slices.Equal(s2, s3) {
			return i
		}
	}

	return -1
}

func performCycle(grid *[][]byte) {
	debug := false
	// shift up, left, down, right
	shiftGrid(grid, 'N')
	if debug {
		for _, l := range *grid {
			fmt.Printf("%s\n", l)
		}
		fmt.Printf("\n")
	}
	shiftGrid(grid, 'W')
	if debug {
		for _, l := range *grid {
			fmt.Printf("%s\n", l)
		}
		fmt.Printf("\n")
	}
	shiftGrid(grid, 'S')
	if debug {
		for _, l := range *grid {
			fmt.Printf("%s\n", l)
		}
		fmt.Printf("\n")
	}
	shiftGrid(grid, 'E')
	if debug {
		for _, l := range *grid {
			fmt.Printf("%s\n", l)
		}
	}
}

func shiftGrid(grid *[][]byte, direction byte) {
	xMove, yMove := 0, 0
	i, jStart := 0, 0
	iIncrement := 1
	jIncrement := 1

	switch direction {
	case 'N':
		yMove = -1
	case 'W':
		xMove = -1
	case 'S':
		yMove = 1
		i = len(*grid) - 1
		iIncrement = -1
	case 'E':
		xMove = 1
		jStart = len((*grid)[0]) - 1
		jIncrement = -1
	}

	for ; i < len(*grid) && i >= 0; i += iIncrement {
		j := jStart
		for ; j < len((*grid)[i]) && j >= 0; j += jIncrement {
			// check to see if the item in the current position is a rock
			// if it is check to see if it can move up, and by how much
			if (*grid)[i][j] == 'O' {
				for k := 1; ; k++ {
					vx := j + (k * xMove)
					vy := i + (k * yMove)

					if vx < 0 || vy < 0 || vx >= len((*grid)[0]) || vy >= len(*grid) {
						break
					}

					if utils.IsPeriod((*grid)[vy][vx]) {
						(*grid)[vy][vx] = 'O'             // move up the rock
						(*grid)[vy-yMove][vx-xMove] = '.' // erase its path
						continue
					}
					break
				}

			}
		}
	}
}

func getGridWeight(grid *[][]byte) int {
	sum := 0
	weight := len(*grid)
	for _, v := range *grid {
		for _, j := range v {
			if j == 'O' {
				sum += weight
			}
		}
		weight--
	}
	return sum
}
