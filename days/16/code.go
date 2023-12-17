package _16

import (
	"advent_of_code23/utils"
	"bytes"
	"fmt"
	"slices"
	"strings"
)

type Coor struct {
	x   int
	y   int
	dir byte // N, E, S, W
}

func Problem1(inputFileName string) int {
	grid := bytes.Split(utils.GetFileContent(inputFileName), []byte("\n"))
	c := make(chan int)
	go getCountForStartingLocation(c, &grid, Coor{0, 0, 'E'})

	sum := 0
	for i := 0; i < 1; i++ {
		sum += <-c
	}

	return sum
}

func Problem2(inputFileName string) int {
	grid := bytes.Split(utils.GetFileContent(inputFileName), []byte("\n"))

	c := make(chan int)
	totalNum := 0
	for i := range grid {
		go getCountForStartingLocation(c, &grid, Coor{0, i, 'E'})
		go getCountForStartingLocation(c, &grid, Coor{len(grid[i]) - 1, i, 'W'})
		totalNum += 2
	}
	for i := range grid[0] {
		go getCountForStartingLocation(c, &grid, Coor{i, 0, 'S'})
		go getCountForStartingLocation(c, &grid, Coor{i, len(grid) - 1, 'N'})
		totalNum += 2
	}
	fmt.Printf("Total number of entry points: %d\n", totalNum)

	m := 0
	for i := 0; i < totalNum; i++ {
		v := <-c
		if v > m {
			m = v
		}
	}

	return m
}

func getCountForStartingLocation(c chan int, grid *[][]byte, coor Coor) {
	history := make(map[string][]string)

	moveThroughGridR(grid, &history, coor)
	c <- len(history)
}

func moveThroughGridR(grid *[][]byte, history *map[string][]string, coor Coor) {
	if coor.x < 0 || coor.x >= len((*grid)[0]) || coor.y < 0 || coor.y >= len(*grid) {
		return
	}
	key := strings.Join([]string{fmt.Sprint(coor.y), fmt.Sprint(coor.x)}, "-")

	if l, ok := (*history)[key]; ok {
		if slices.Contains(l, string(coor.dir)) {
			return
		}
		(*history)[key] = append(l, string(coor.dir))
	} else {
		(*history)[key] = []string{string(coor.dir)}
	}

	for _, move := range getMoves(grid, coor) {
		moveThroughGridR(grid, history, move)
	}
}

func getMoves(grid *[][]byte, coor Coor) []Coor {
	switch (*grid)[coor.y][coor.x] {
	case '.':
		return []Coor{getCoorFromDirection(coor.x, coor.y, coor.dir)}
	case '|':
		if coor.dir == 'N' || coor.dir == 'S' {
			return []Coor{getCoorFromDirection(coor.x, coor.y, coor.dir)}
		}
		return []Coor{
			getCoorFromDirection(coor.x, coor.y, 'N'),
			getCoorFromDirection(coor.x, coor.y, 'S'),
		}
	case '-':
		if coor.dir == 'E' || coor.dir == 'W' {
			return []Coor{getCoorFromDirection(coor.x, coor.y, coor.dir)}
		}
		return []Coor{
			getCoorFromDirection(coor.x, coor.y, 'E'),
			getCoorFromDirection(coor.x, coor.y, 'W'),
		}
	case '\\':
		switch coor.dir {
		case 'N':
			return []Coor{getCoorFromDirection(coor.x, coor.y, 'W')}
		case 'E':
			return []Coor{getCoorFromDirection(coor.x, coor.y, 'S')}
		case 'S':
			return []Coor{getCoorFromDirection(coor.x, coor.y, 'E')}
		case 'W':
			return []Coor{getCoorFromDirection(coor.x, coor.y, 'N')}
		}
		panic("Invalid move")
	case '/':
		switch coor.dir {
		case 'N':
			return []Coor{getCoorFromDirection(coor.x, coor.y, 'E')}
		case 'E':
			return []Coor{getCoorFromDirection(coor.x, coor.y, 'N')}
		case 'S':
			return []Coor{getCoorFromDirection(coor.x, coor.y, 'W')}
		case 'W':
			return []Coor{getCoorFromDirection(coor.x, coor.y, 'S')}
		}
		panic("Invalid move")
	}
	panic("Invalid move")
}

func getCoorFromDirection(x, y int, dir byte) Coor {
	switch dir {
	case 'N':
		return Coor{x, y - 1, dir}
	case 'E':
		return Coor{x + 1, y, dir}
	case 'S':
		return Coor{x, y + 1, dir}
	case 'W':
		return Coor{x - 1, y, dir}
	}
	panic("Invalid direction")
}
