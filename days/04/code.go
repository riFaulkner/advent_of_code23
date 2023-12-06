package _4

import (
	"advent_of_code23/utils"
	"fmt"
)

type ScratchCard struct {
	winningNumbers []int
	numbers        []int
}

func GetTotalScratchCards(inputFileName string) int {
	content := utils.GetFileContent(inputFileName)
	cardStrings := utils.SplitByteArrayByLine(content)

	cards := serializeScratchCards(cardStrings)

	// Card tracker will count how many of each card we have, where k is the card number and v is the count of that card we have,
	//including copies. If we find a winner we'll add the new cards to the map.
	cardTracker := generateCardMap(len(cards))

	for i, card := range cards {
		numMatches := getNumberCardMatches(card)
		numCurrentCard := cardTracker[i]
		for j := 1; j < numMatches+1; j++ {
			if i+j >= len(cards) {
				break
			}
			cardTracker[i+j] += numCurrentCard
		}
	}

	sum := calculateNumCards(cardTracker)

	fmt.Printf("Total: %d\n", sum)
	return sum
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

func calculateNumCards(cardTracker map[int]int) int {
	sum := 0
	for _, v := range cardTracker {
		sum += v
	}
	return sum
}

func generateCardMap(size int) map[int]int {
	cardMap := make(map[int]int, size)
	for i := 0; i < size; i++ {
		cardMap[i] = 1
	}

	return cardMap
}

func serializeScratchCards(cardStrings [][]byte) []ScratchCard {
	cards := make([]ScratchCard, 0)
	for _, cardString := range cardStrings {
		cards = append(cards, serializeScratchCard(cardString))
	}

	return cards
}

func getNumberCardMatches(card ScratchCard) int {
	matches := 0
	for _, winningNumber := range card.winningNumbers {
		for _, number := range card.numbers {
			if number == winningNumber {
				matches++
			}
		}
	}
	return matches
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
