package _13

import (
	"advent_of_code23/utils"
	"bytes"
)

func Problem1(inputFileName string, failures int) int {
	grids := bytes.Split(utils.GetFileContent(inputFileName), []byte("\n\n"))

	sum := 0
	for _, grid := range grids {
		g := bytes.Split(grid, []byte("\n"))
		sum += getHorizontalSplitRows(g, failures) * 100
		sum += getVerticalSplitColumns(g, failures)
	}

	return sum
}

func getHorizontalSplitRows(grid [][]byte, failures int) int {
	for i := 1; i < len(grid); i++ {
		doMatch, _ := doRowsMatch(&grid, i, i-1, failures)
		if doMatch {
			f := failures
			for j := 0; j <= i; j++ {
				if i <= j || i+j > len(grid)-1 { // If the rows are unbalanced it doesn't matter that the end cant continue as long as the beginning is correct
					if f == 0 {
						return i
					}
					break
				}
				m, rF := doRowsMatch(&grid, i+j, i-1-j, f)
				if !m {
					break
				}
				f = f - rF
			}
		}
	}

	return 0
}

func getVerticalSplitColumns(grid [][]byte, failures int) int {
	for i := 1; i < len(grid[0]); i++ {
		doMatch, _ := doColumnsMatch(&grid, i, i-1, failures)
		if doMatch {
			f := failures
			for j := 0; j <= i; j++ {
				if j == i || i+j > len(grid[0])-1 {
					if f == 0 {
						return i
					}
					break
				}

				m, rF := doColumnsMatch(&grid, i+j, i-1-j, f)
				if !m {
					break
				}
				f = f - rF
			}
		}
	}

	return 0
}

func doColumnsMatch(grid *[][]byte, cA, cB, failures int) (bool, int) {
	f := 0
	for _, row := range *grid {
		if cA > len(row) || cB > len(row) {
			return false, 0
		}
		if row[cA] != row[cB] {
			if failures-f == 0 {
				return false, 0
			}
			f += 1
		}
	}
	return true, f
}

func doRowsMatch(grid *[][]byte, rA, rB, failures int) (bool, int) {
	f := 0
	for i := 0; i < len((*grid)[rA]); i++ {
		if (*grid)[rA][i] != (*grid)[rB][i] {
			if failures-f == 0 {
				return false, 0
			}
			f += 1
		}
	}
	return true, f
}
