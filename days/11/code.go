package _11

import (
	"advent_of_code23/utils"
	"fmt"
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

	// then Generate all the pairs of Galaxies, and calculate the distance between them
	sum := 0
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			sum += getDistance(galaxies[i], galaxies[j])
		}
	}

	// return the sum of the distances
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

func expandGalaxy(grid *[][]byte) {
	totalRows := len(*grid)
	for i := 0; i < totalRows; i++ {
		if !rowHasGalaxy((*grid)[i]) {
			*grid = append((*grid)[:i+1], (*grid)[i:]...)
			i++
			totalRows++
			for j := 0; j < len((*grid)[i]); j++ {
				(*grid)[i][j] = '.'
			}
		}
	}
	totalColumns := len((*grid)[0])
	for i := 0; i < totalColumns; i++ {
		if !columHasGalaxy(*grid, i) {
			fmt.Printf("\n\n\n")
			for j := 0; j < len(*grid); j++ {
				newRow := append((*grid)[j][:i+1], (*grid)[j][i:]...)
				newRow[i] = '.'
				(*grid)[j] = newRow
				fmt.Printf("grid: %v\n", (*grid)[j])
			}
			i++
			totalColumns++
		}
	}
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
