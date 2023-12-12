package _0

import (
	"advent_of_code23/utils"
	"fmt"
	"slices"
)

type Coor struct {
	x, y int
}

var moveMap = map[byte][]string{
	'|': {"north", "south"},
	'-': {"east", "west"},
	'L': {"north", "east"},
	'J': {"north", "west"},
	'7': {"south", "west"},
	'F': {"south", "east"},
	'.': {},
}

//| is a vertical pipe connecting north and south.
//- is a horizontal pipe connecting east and west.
//L is a 90-degree bend connecting north and east.
//J is a 90-degree bend connecting north and west.
//7 is a 90-degree bend connecting south and west.
//F is a 90-degree bend connecting south and east.
//. is ground; there is no pipe in this tile.
//S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.

func Problem1(inputFileName string) int {
	content := utils.SplitByteArrayByLine(utils.GetFileContent(inputFileName))
	// now should have a nice lil grid. Find the 'S' in the grid as our starting point
	startingCoor := getStartingPoint(content)
	fmt.Printf("Starting point is %v\n", startingCoor)

	startMoves, directions := getStartingPointPossibleMoves(content, startingCoor)
	fmt.Printf("Starting moves are %v, and directions%v\n", startMoves, directions)

	return findLastR(startingCoor, startingCoor, directions[0], 0, &content) / 2
}

func findLastR(start, end Coor, direction string, count int, grid *[][]byte) int {
	newCoor := makeMove(start, direction)
	if newCoor.x == end.x && newCoor.y == end.y {
		return count + 1
	}
	newCoorMoves := moveMap[(*grid)[newCoor.y][newCoor.x]]
	move := newCoorMoves[0]
	if move == getDirectionInverse(direction) {
		move = newCoorMoves[1]
	}
	return findLastR(newCoor, end, move, count+1, grid)
}

// Make move in a direction but does not check if it is valid
// if we go out of bounds, we panic
func makeMove(start Coor, direction string) Coor {
	switch direction {
	case "north":
		return Coor{
			x: start.x,
			y: start.y - 1,
		}
	case "south":
		return Coor{
			x: start.x,
			y: start.y + 1,
		}
	case "east":
		return Coor{
			x: start.x + 1,
			y: start.y,
		}
	case "west":
		return Coor{
			x: start.x - 1,
			y: start.y,
		}
	default:
		panic("Invalid direction")
	}
}

func getStartingPoint(grid [][]byte) Coor {
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 'S' {
				return Coor{
					x: j,
					y: i,
				}
			}
		}
	}
	panic("No starting point found")
}

func getStartingPointPossibleMoves(grid [][]byte, startingCoor Coor) ([]Coor, []string) {
	coor := make([]Coor, 0)
	directions := make([]string, 0)
	i := startingCoor.y
	j := startingCoor.x

	// check up
	if i > 0 {
		moves := moveMap[grid[i-1][j]]
		if slices.Contains(moves, "south") {
			coor = append(coor, Coor{i - 1, j})
			directions = append(directions, "north")
		}
	}
	// check down
	if i < len(grid)-1 {
		moves := moveMap[grid[i+1][j]]
		if slices.Contains(moves, "north") {
			coor = append(coor, Coor{i + 1, j})
			directions = append(directions, "south")
		}

	}
	// check left
	if j > 0 {
		moves := moveMap[grid[i][j-1]]
		if slices.Contains(moves, "east") {
			coor = append(coor, Coor{i, j - 1})
			directions = append(directions, "west")
		}
	}

	// check right
	if j < len(grid[i])-1 {
		moves := moveMap[grid[i][j+1]]
		if slices.Contains(moves, "west") {
			coor = append(coor, Coor{i, j + 1})
			directions = append(directions, "east")
		}
	}

	return coor, directions
}

func getDirectionInverse(direction string) string {
	switch direction {
	case "north":
		return "south"
	case "south":
		return "north"
	case "east":
		return "west"
	case "west":
		return "east"
	default:
		panic("Invalid direction")
	}
}
