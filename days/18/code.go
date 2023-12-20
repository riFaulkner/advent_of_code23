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

func Problem1(inputFileName string, part2 bool) int {
	lines := strings.Split(utils.GetFileAsString(inputFileName), "\n")

	maxC, minC := 0, 0 // max and min columns from the starting position
	maxR, minR := 0, 0 // max and min rows from the starting position

	cC, cR := 0, 0 // current column and row
	inst := make([]instruction, len(lines))
	for i := range lines {
		ins := serializeInstruction(&lines[i], part2)
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

	coorSet := orderedmap.New()
	gR := 1 + maxR - minR
	gC := 1 + maxC - minC
	cR, cC = gR-1, gC-1

	cS := []int{cC + 1, gR - cR}
	coorSet.Set(fmt.Sprintf("%v,%v", cC+1, gR-cR), cS)

	o := 0
	for i := range inst {
		ins := inst[i]
		moveCoor := directions[ins.direction]

		// might need to -1
		o += ins.distance

		cR = cR + moveCoor[0]*ins.distance
		cC = cC + moveCoor[1]*ins.distance

		// Inverting the coords to get a more traditional grid.
		cS = []int{cC + 1, gR - cR}
		coorSet.Set(fmt.Sprintf("%v,%v", cC+1, gR-cR), cS)
	}

	a := calcAreaFromCoorSet(coorSet)

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

func serializeInstruction(line *string, part2 bool) instruction {
	seg := strings.Split(*line, " ")
	n, err := strconv.Atoi(seg[1])
	if err != nil {
		panic(err)
	}
	dir := 0
	switch seg[0][0] {
	case 'R':
		dir = 0
	case 'D':
		dir = 1
	case 'L':
		dir = 2
	case 'U':
		dir = 3
	}

	if part2 {
		var ok error
		if dir, ok = strconv.Atoi(string(seg[2][len(seg[2])-2])); ok != nil {
			panic(ok)
		}
		s := seg[2][2 : len(seg[2])-2] // Get the number without the last item
		hN, err := strconv.ParseInt(s, 16, 64)
		if err != nil {
			panic(err)
		}
		n = int(hN)
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
