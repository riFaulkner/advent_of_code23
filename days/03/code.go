package _3

import (
	"advent_of_code23/utils"
	"fmt"
	"slices"
)

type Coordinate struct {
	x int
	y int
}

func GetSumOfSchematicParts(inputFileName string) int {
	content := utils.GetFileContent(inputFileName)
	schematicGrid := utils.SplitByteArrayByLine(content)

	runningPartTotal := 0

	// Next iterate though the grid, if we find a digit then we want to keep track of the starting index.
	// once we get the whole number, we want to create a list of all coordinates that surround the number.
	// next safely check all those coordinates see if any of them contain a symbol that is not a period
	// if they do, then we want to add that number to the total.

	for i, row := range schematicGrid {
		partNumberBytes := make([]byte, 0)
		xStart := 0
		for j, x := range row {
			if utils.IsDigit(x) {
				if len(partNumberBytes) == 0 {
					xStart = j
				}
				partNumberBytes = append(partNumberBytes, x)
				continue
			} else {
				if len(partNumberBytes) > 0 {
					partNumber := utils.GetNumberFromBytes(partNumberBytes)
					partNumberBytes = make([]byte, 0)
					// Now we have the part number, we need to get the coordinates that surround it.
					// We need to check the following coordinates
					surroundingCoordinates := generateSurroundingCoordinates(xStart, i, j-1)
					if doSurroundingCoordinatesContainSymbol(schematicGrid, surroundingCoordinates) {
						runningPartTotal += partNumber
						fmt.Printf("Adding %d , running total is now %d\n", partNumber, runningPartTotal)
					}
				}
			}

		}
		if len(partNumberBytes) > 0 {
			partNumber := utils.GetNumberFromBytes(partNumberBytes)
			partNumberBytes = make([]byte, 0)
			// Now we have the part number, we need to get the coordinates that surround it.
			// We need to check the following coordinates
			surroundingCoordinates := generateSurroundingCoordinates(xStart, i, len(row)-1)
			if doSurroundingCoordinatesContainSymbol(schematicGrid, surroundingCoordinates) {
				runningPartTotal += partNumber
				fmt.Printf("Adding %d , running total is now %d\n", partNumber, runningPartTotal)
			}
		}
	}

	return runningPartTotal
}

func GetSumOfGearRatios(inputFileName string) int {
	content := utils.GetFileContent(inputFileName)
	schematicGrid := utils.SplitByteArrayByLine(content)
	runningTotal := 0

	// Iterate through the schematic grid, if we find a potential gear indicator '*' then we want to check the surrounding coordinates
	// first generate the surrounding coordinates, then check each one to see if it is a digit, if it is then we watnt to get the whole number associated with it.
	// I need a way to remove dupes. e.g. if a single number has multiple digits inside the gear indicator's radius, then we only want to count it once.
	// If the indicator has two numbers in its range, it is considered valid, multiply those together.
	// If more than 2 numbers are in the range, then it is invalid and we should not count it.

	for i, row := range schematicGrid {
		for j, b := range row {
			if utils.IsAstrix(b) {
				surroundingCoordinates := generateSurroundingCoordinates(j, i, j)
				numbers := findNumbersInCoordinates(schematicGrid, surroundingCoordinates)
				if len(numbers) == 2 {
					runningTotal += numbers[0] * numbers[1]
					fmt.Printf("Adding %d , running total is now %d\n", numbers[0]*numbers[1], runningTotal)
				}
				continue
			}
		}
	}

	return runningTotal
}

func generateSurroundingCoordinates(x, y, x2 int) []Coordinate {
	coordinates := make([]Coordinate, 0)
	for i := x - 1; i <= x2+1; i++ {
		coordinates = append(coordinates, Coordinate{i, y - 1})
		coordinates = append(coordinates, Coordinate{i, y + 1})
	}
	coordinates = append(coordinates, Coordinate{x - 1, y})
	coordinates = append(coordinates, Coordinate{x2 + 1, y})
	return coordinates
}

func doSurroundingCoordinatesContainSymbol(schematicGrid [][]byte, coordinates []Coordinate) bool {
	min := 0
	maxX := len(schematicGrid[0]) - 1
	maxY := len(schematicGrid) - 1

	for _, coordinate := range coordinates {
		if coordinate.x < min || coordinate.x > maxX || coordinate.y < min || coordinate.y > maxY {
			continue
		}
		symbol := schematicGrid[coordinate.y][coordinate.x]
		if !utils.IsPeriod(symbol) && !utils.IsDigit(symbol) {
			return true
		}
	}
	return false
}

func findNumbersInCoordinates(schematicGrid [][]byte, coordinates []Coordinate) []int {
	numbers := make([]int, 0)
	// create a map of coordinates where the y is the key and there is a list of x values
	excludedCoordinates := make(map[int][]int)

	for _, coordinate := range coordinates {
		symbol := schematicGrid[coordinate.y][coordinate.x]
		if utils.IsDigit(symbol) && !slices.Contains(excludedCoordinates[coordinate.y], coordinate.x) {
			// Find the whole number
			numberBytes := append(getPriorDigit(schematicGrid, coordinate.x, coordinate.y), symbol)
			// Why make a new one? So we can correctly block the used coordinates below before concatenating the next digits
			nextDigits := getNextDigit(schematicGrid, coordinate.x, coordinate.y, make([]byte, 0))
			numberBytes = append(numberBytes, nextDigits...)
			for i, _ := range nextDigits {
				xList := excludedCoordinates[coordinate.y]
				excludedCoordinates[coordinate.y] = append(xList, coordinate.x+i+1)
			}

			numbers = append(numbers, utils.GetNumberFromBytes(numberBytes))

		}
	}
	return numbers
}

func getPriorDigit(schematicGrid [][]byte, x, y int) []byte {
	if x-1 < 0 || !utils.IsDigit(schematicGrid[y][x-1]) {
		return make([]byte, 0)
	}
	return append(getPriorDigit(schematicGrid, x-1, y), schematicGrid[y][x-1])
}

func getNextDigit(schematicGrid [][]byte, x, y int, numberBytes []byte) []byte {
	if x+1 >= len(schematicGrid[y]) || !utils.IsDigit(schematicGrid[y][x+1]) {
		return numberBytes
	}
	return getNextDigit(schematicGrid, x+1, y, append(numberBytes, schematicGrid[y][x+1]))
}
