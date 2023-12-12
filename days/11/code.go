package _11

import (
	"advent_of_code23/utils"
	"math"
)

type Coor struct {
	x int
	y int
}

func Problem1(inputFileName string, expansionMultiplier int) int {
	content := utils.SplitByteArrayByLine(utils.GetFileContent(inputFileName))

	galaxies := findAllGalaxies(content)
	cExpand := getAllColumnsToExpand(content)
	rExpand := getAllRowsToExpand(content)

	for i, g := range galaxies {
		for c := range cExpand {
			if g.y > cExpand[len(cExpand)-1-c] {
				galaxies[i].y = galaxies[i].y + expansionMultiplier
			}
		}

		for r := range rExpand {
			if g.x > rExpand[len(rExpand)-1-r] {
				galaxies[i].x = galaxies[i].x + expansionMultiplier
			}
		}
	}

	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += getDistance(galaxies[i], galaxies[j])
		}
	}

	return sum
}

func getDistance(coor1, coor2 Coor) int {
	return int(math.Abs(float64(coor1.x-coor2.x)) + math.Abs(float64(coor1.y-coor2.y)))
}

func findAllGalaxies(grid [][]byte) []Coor {
	var galaxies []Coor
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '#' {
				galaxies = append(galaxies, Coor{i, j})
			}
		}
	}
	return galaxies
}

func getAllColumnsToExpand(grid [][]byte) []int {
	var columns []int
	for i := 0; i < len(grid); i++ {
		if !columHasGalaxy(grid, i) {
			columns = append(columns, i)
		}
	}
	return columns
}

func columHasGalaxy(grid [][]byte, index int) bool {
	for i := 0; i < len(grid); i++ {
		if len(grid[i]) < index {
			continue
		}
		if grid[i][index] == '#' {
			return true
		}
	}
	return false
}

func getAllRowsToExpand(grid [][]byte) []int {
	var rows []int
	for i := 0; i < len(grid); i++ {
		if !rowHasGalaxy(grid[i]) {
			rows = append(rows, i)
		}
	}
	return rows
}

func rowHasGalaxy(row []byte) bool {
	for i := 0; i < len(row); i++ {
		if row[i] == '#' {
			return true
		}
	}
	return false
}
