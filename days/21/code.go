package _21

import (
	"advent_of_code23/utils"
	"bytes"
	"fmt"
)

func Problem1(inputFileName string, numSteps int, part2 bool) int {
	grid := bytes.Split(utils.GetFileContent(inputFileName), []byte("\n"))
	startR, startC := findStartingPosition(&grid)

	visited, destinations := make(map[string]bool), make(map[string]bool)

	// step through and see where you can move.
	// we may be able to do this recursively
	takeStepR(&grid, startR, startC, 0, numSteps, &visited, &destinations, part2)

	return len(destinations)
}

func takeStepR(grid *[][]byte, r, c, count, cap int, visited, destinations *map[string]bool, part2 bool) {
	key := fmt.Sprintf("%d-%d-%d", r, c, count)
	if _, ok := (*visited)[key]; ok {
		return
	}
	(*visited)[key] = true
	if len(*visited)%1_000_000 == 0 {
		fmt.Printf("visited: %d, destinations: %d\n", len(*visited), len(*destinations))
	}

	if count >= cap { // Termination condition
		(*destinations)[key] = true

		return
	}

	// Find my neighbors based on my current position. Up down left right.
	for _, direction := range utils.Directions {
		newR, newC := r+direction[0], c+direction[1]
		if !part2 && !utils.IsInbounds(grid, newR, newC) || (*grid)[newR][newC] == '#' {
			continue
		}
		takeStepR(grid, newR, newC, count+1, cap, visited, destinations, part2)
	}
}

func getGridPositions(grid *[][]byte, r, c int, direction []int) (int, int) {

	newR, newC := r+direction[0]%len(*grid), c+direction[1]%len((*grid)[0])

	return newR, newC
}

func findStartingPosition(grid *[][]byte) (int, int) {
	for r, row := range *grid {
		for c, cell := range row {
			if cell == 'S' {
				return r, c
			}
		}
	}
	panic("No starting position found")
}
