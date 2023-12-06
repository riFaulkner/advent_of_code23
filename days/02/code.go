package _2

import (
	"advent_of_code23/utils"
	"fmt"
)

type Game struct {
	pointValue int
	rounds     []Round
}

type Round struct {
	red   int
	green int
	blue  int
}

func GetSumOfPossibleGames(inputFileName string, roundMaxes Round) int {
	content := utils.GetFileContent(inputFileName)
	games := serializeGames(utils.SplitByteArrayByLine(content))
	sum := 0

	for _, game := range games {
		isGamePossible := true
		// only 12 red cubes, 13 green cubes, and 14 blue cubes
		for _, round := range game.rounds {
			if round.red > roundMaxes.red || round.green > roundMaxes.green || round.blue > roundMaxes.blue {
				isGamePossible = false
				break
			}
		}
		if isGamePossible {
			sum += game.pointValue
		}
	}

	fmt.Printf("Total: %d\n", len(games))
	return sum
}

func GetPowerOfMinCubesPossible(inputFileName string) int {
	content := utils.GetFileContent(inputFileName)
	games := serializeGames(utils.SplitByteArrayByLine(content))
	sum := 0

	for _, game := range games {
		roundMaxes := Round{
			red:   0,
			green: 0,
			blue:  0,
		}
		for _, round := range game.rounds {
			if round.red > roundMaxes.red {
				roundMaxes.red = round.red
			}
			if round.green > roundMaxes.green {
				roundMaxes.green = round.green
			}
			if round.blue > roundMaxes.blue {
				roundMaxes.blue = round.blue
			}
		}
		sum += roundMaxes.red * roundMaxes.green * roundMaxes.blue
	}

	return sum
}

func serializeGames(gameStrings [][]byte) []Game {
	games := make([]Game, len(gameStrings))
	for i, game := range gameStrings {
		games[i] = getGameFromBytes(game)
	}

	return games
}

func getGameFromBytes(b []byte) Game {
	// Game is shaped like this: Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
	// where the first number is the point value, and the rounds are seperated by semicolons
	pointValue := extractPointValueFromBytes(&b)
	rounds := extractRoundsFromBytes(&b)

	return Game{
		pointValue: pointValue,
		rounds:     rounds,
	}
}

func extractPointValueFromBytes(b *[]byte) int {
	// find first number in the array of bytes
	numberStartIndex := 0
	numberBytes := make([]byte, 0)
	for i, v := range *b {
		if utils.IsDigit(v) {
			if numberStartIndex == 0 {
				numberStartIndex = i
			}
			numberBytes = append(numberBytes, v)
		} else {
			if numberStartIndex != 0 {
				newB := (*b)[i+1:]
				*b = newB
				break
			}
		}
	}

	return utils.GetNumberFromBytes(numberBytes)
}

func extractRoundsFromBytes(b *[]byte) []Round {
	rounds := make([]Round, 0)
	for len(*b) > 0 {
		round, err := extractNextRoundFromBytes(b)
		if err != nil {
			break
		}
		rounds = append(rounds, round)
	}

	return rounds
}

func extractNextRoundFromBytes(b *[]byte) (Round, error) {
	// Each round is shaped like this: 3 blue, 4 red;
	// where the first number is the number of balls, and the color is the color of the balls
	// and the rounds are seperated by semicolons
	round := Round{
		red:   0,
		green: 0,
		blue:  0,
	}

	numberBytes := make([]byte, 0)
	colorBytes := make([]byte, 0)

	for i, v := range *b {
		if utils.IsDigit(v) {
			numberBytes = append(numberBytes, v)
		}

		if utils.IsLetter(v) {
			colorBytes = append(colorBytes, v)
		}

		if utils.IsComma(v) {
			// Evaluate the color and add the number of balls to the round to that color
			ballCount := utils.GetNumberFromBytes(numberBytes)
			switch string(colorBytes) {
			case "red":
				round.red = ballCount
			case "green":
				round.green = ballCount
			case "blue":
				round.blue = ballCount
			default:
				return Round{}, fmt.Errorf("unknown color: %s", string(colorBytes))
			}
			numberBytes = make([]byte, 0)
			colorBytes = make([]byte, 0)
		}

		// end of round
		if utils.IsSemiColon(v) {
			// Evaluate the color and add the number of balls to the round to that color
			ballCount := utils.GetNumberFromBytes(numberBytes)
			switch string(colorBytes) {
			case "red":
				round.red = ballCount
			case "green":
				round.green = ballCount
			case "blue":
				round.blue = ballCount
			default:
				return Round{}, fmt.Errorf("unknown color: %s", string(colorBytes))
			}
			numberBytes = make([]byte, 0)
			colorBytes = make([]byte, 0)

			// Remove this round from the line
			*b = (*b)[i+1:]
			return round, nil
		}
	}
	// Remove this round from the line
	// Evaluate the color and add the number of balls to the round to that color
	ballCount := utils.GetNumberFromBytes(numberBytes)
	switch string(colorBytes) {
	case "red":
		round.red = ballCount
	case "green":
		round.green = ballCount
	case "blue":
		round.blue = ballCount
	default:
		return Round{}, fmt.Errorf("unknown color: %s", string(colorBytes))
	}
	numberBytes = make([]byte, 0)
	colorBytes = make([]byte, 0)
	*b = (*b)[len(*b):]

	// Each round can have the colors red, green and blue find out how many if any, of each color there are
	return round, nil
}
