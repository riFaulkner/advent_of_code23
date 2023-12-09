package _6

import (
	"advent_of_code23/utils"
	"regexp"
	"strconv"
	"strings"
)

func GetPowerOfWaysToWin(inputFile string) int {
	content := utils.GetFileContent(inputFile)
	sections := regexp.MustCompile(`[\d]+`).FindAllString(string(content), -1)

	midwayPoint := len(sections) / 2
	winOptions := 0

	for i := 0; i < midwayPoint; i++ {
		duration, err := strconv.Atoi(sections[i])
		if err != nil {
			panic(err)
		}
		best, err := strconv.Atoi(sections[i+midwayPoint])
		if err != nil {
			panic(err)
		}

		winOptionsForRace := calculateRaceWinOptions(duration, best)
		if winOptions == 0 {
			winOptions = winOptionsForRace
		} else {
			winOptions = winOptions * winOptionsForRace
		}
	}
	return winOptions
}

func ProblemTwo(inputFile string) int {
	content := utils.GetFileContent(inputFile)

	lines := strings.Split(strings.ReplaceAll(string(content), " ", ""), "\n")
	d := regexp.MustCompile(`[\d]+`).FindAllString(lines[0], -1)
	duration, err := strconv.Atoi(d[0])
	if err != nil {
		panic(err)
	}
	b := regexp.MustCompile(`[\d]+`).FindAllString(lines[1], -1)
	best, err := strconv.Atoi(b[0])
	if err != nil {
		panic(err)
	}

	return calculateRaceWinOptions(duration, best)
}

func calculateRaceWinOptions(duration, best int) int {
	winOptions := 0
	for i := 0; i < duration; i++ {
		distance := i * (duration - i)
		if best < distance {
			winOptions++
		}
	}
	return winOptions
}
