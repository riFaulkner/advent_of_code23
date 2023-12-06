package _4

import (
	"advent_of_code23/utils"
	"fmt"
)

type ScratchCard struct {
	winningNumbers []int
	numbers        []int
}

func GetTotalScratchCardPoints(inputFileName string) int {
	content := utils.GetFileContent(inputFileName)
	cardStrings := utils.SplitByteArrayByLine(content)

	cards := serializeScratchCards(cardStrings)

	sum := 0
	for _, card := range cards {
		sum += getCardPoints(card)
	}

	fmt.Printf("Total: %d\n", sum)
	return sum
}

func serializeScratchCards(cardStrings [][]byte) []ScratchCard {
	cards := make([]ScratchCard, 0)
	for _, cardString := range cardStrings {
		cards = append(cards, serializeScratchCard(cardString))
	}

	return cards
}

func getCardPoints(card ScratchCard) int {
	points := 0
	for _, winningNumber := range card.winningNumbers {
		for _, number := range card.numbers {
			if number == winningNumber {
				if points == 0 {
					points++
				} else {
					points += points
				}
			}
		}
	}

	return points
}

func serializeScratchCard(cardBytes []byte) ScratchCard {
	winningNumbers := make([]int, 0)
	numbers := make([]int, 0)

	startingIndex := 0
	numberBytes := make([]byte, 0)
	workingList := &winningNumbers
	for i, b := range cardBytes {
		// Card string is in this format: Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
		// first get to the : starting index lets us skip other characters until we're at the start of the card
		if startingIndex == 0 {
			if utils.IsColon(b) {
				startingIndex = i + 1
			}
			continue
		}
		if utils.IsDigit(b) {
			numberBytes = append(numberBytes, b)
			continue
		}
		if utils.IsBar(b) {
			winningNumbers = *workingList
			workingList = &numbers
			continue
		}
		if len(numberBytes) > 0 {
			// TODO: better way to do this?
			newList := append(*workingList, utils.GetNumberFromBytes(numberBytes))
			numberBytes = make([]byte, 0)
			workingList = &newList
		}
	}

	if len(numberBytes) > 0 {
		// TODO: better way to do this?
		newList := append(*workingList, utils.GetNumberFromBytes(numberBytes))
		numberBytes = make([]byte, 0)
		workingList = &newList
	}
	numbers = *workingList
	return ScratchCard{
		winningNumbers: winningNumbers,
		numbers:        numbers,
	}
}
