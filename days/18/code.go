package _18

import (
	"advent_of_code23/utils"
	"fmt"
	orderedmap "github.com/wk8/go-ordered-map"
	"math"
	"strconv"
	"strings"
)

type instruction struct {
	direction int
	distance  int
	color     string
}

var directions = [][]int{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func Problem1(inputFileName string) int {
	debug := true
	lines := strings.Split(utils.GetFileAsString(inputFileName), "\n")

	maxC, minC := 0, 0 // max and min columns from the starting position
	maxR, minR := 0, 0 // max and min rows from the starting position

	cC, cR := 0, 0 // current column and row
	inst := make([]instruction, len(lines))
	for i := range lines {
		ins := serializeInstruction(&lines[i])
		inst[i] = ins
		moveCoor := directions[ins.direction]
		cR += moveCoor[0] * ins.distance
		cC += moveCoor[1] * ins.distance

		// If we're on a new row or column, update the max and min
		if cC > maxC {
			maxC = cC
		}
		if cC < minC {
			minC = cC
		}
		if cR > maxR {
			maxR = cR
		}
		if cR < minR {
			minR = cR
		}
	}

	// Now we have the size of the grid, lets create it
	grid := make([][]string, 1+maxR-minR)
	r, c := 0, 0

	for i := minR; i <= maxR; i++ {
		grid[r] = make([]string, 1+maxC-minC)
		for j := minC; j <= maxC; j++ {
			grid[r][c] = "."
			c++
		}
		r++
		c = 0
	}

	if debug {
		printGrid(&grid)
	}

	coorSet := orderedmap.New()
	cR, cC = len(grid)-1-maxR, len(grid[0])-1-maxC
	rM, cM := len(grid), len(grid[0])
	fmt.Printf("Start at %v,%v\n", rM, cM)

	cS := []int{cC + 1, rM - cR}
	coorSet.Set(fmt.Sprintf("%v,%v", cC+1, rM-cR), cS)

	fmt.Printf("Start at %v,%v\n", cC, cR)
	o := 0
	for i := range inst {
		ins := inst[i]
		moveCoor := directions[ins.direction]

		// for each step in the distance, color the grid
		for j := 0; j <= ins.distance; j++ {
			rMove := moveCoor[0] * j
			cMove := moveCoor[1] * j
			if grid[cR+rMove][cC+cMove] == "." {
				o++
			}

			grid[cR+rMove][cC+cMove] = "#"
		}

		cR = cR + moveCoor[0]*ins.distance
		cC = cC + moveCoor[1]*ins.distance

		// Inverting the coords to get a more traditional grid.
		cS = []int{cC + 1, rM - cR}
		coorSet.Set(fmt.Sprintf("%v,%v", cC+1, rM-cR), cS)

		if debug && false {
			printGrid(&grid)
		}
	}

	a := calcAreaFromCoorSet(coorSet)
	fmt.Printf("Area calcuated: %v\n", a)
	printGrid(&grid)

	// Use Pick's theorem to calculate the around?
	// Not fully sure how this works but we use the
	// area inside and then pick's theorem to calculate the area around
	return a + (o / 2) + 1
}

func calcAreaFromCoorSet(coorSet *orderedmap.OrderedMap) int {
	// multiple diagonals where x, y multiplies x * x+1 and y * y+1
	// this needs to wrap all the way around back to the beginning
	xTop := 0
	yTop := 0
	for item := coorSet.Oldest(); item != nil; item = item.Next() {
		n := item.Next()
		if n == nil {
			n = coorSet.Oldest()
		}
		xTop += (item.Value.([]int)[0] * n.Value.([]int)[1])
		yTop += (item.Value.([]int)[1] * n.Value.([]int)[0])
	}
	t := math.Abs(float64(xTop - yTop))

	a := t / 2
	return int(a)
}

func serializeInstruction(line *string) instruction {
	seg := strings.Split(*line, " ")
	n, err := strconv.Atoi(seg[1])
	if err != nil {
		panic(err)
	}
	dir := 0
	switch seg[0][0] {
	case 'U':
		dir = 0
	case 'R':
		dir = 1
	case 'D':
		dir = 2
	case 'L':
		dir = 3
	}
	return instruction{
		direction: dir,
		distance:  n,
		color:     seg[2],
	}
}

func printGrid(grid *[][]string) {
	for i := range *grid {
		fmt.Printf("%v\n", (*grid)[i])
	}
	fmt.Printf("\n\n")
}
