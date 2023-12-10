package _7

import (
	"advent_of_code23/utils"
	"cmp"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	cards []int
	bid   int
	rank  int
}

const (
	HighCard     = 1
	OnePair      = 2
	TwoPair      = 3
	ThreeOfAKind = 4
	FullHouse    = 5
	FourOfAKind  = 6
	FiveOfAKind  = 7
)

func Problem1(inputFile string, hasJokers bool) int {
	content := strings.Split(utils.GetFileAsString(inputFile), "\n")
	hands := serializeHands(&content, hasJokers)

	slices.SortFunc(hands, func(i, j Hand) int {
		if i.rank == j.rank {
			for c := range i.cards {
				if i.cards[c] == j.cards[c] {
					continue
				}
				return cmp.Compare(i.cards[c], j.cards[c])
			}
			fmt.Printf("TWO IDENDICAL MATCHES: %v %v\n", i, j)
			return 0
		}
		return cmp.Compare(i.rank, j.rank)
	})

	sum := 0
	for i, c := range hands {
		sum += c.bid * (i + 1)
	}
	return sum
}

func serializeHands(content *[]string, hasJokers bool) []Hand {
	hands := make([]Hand, len(*content))
	for i, s := range *content {
		p := strings.Split(s, " ")
		b, err := strconv.Atoi(p[1])
		if err != nil {
			panic(err)
		}

		cards := make([]int, 5)
		for i := 0; i < 5; i++ {
			byteVal := p[0][i]
			if utils.IsDigit(byteVal) {
				v, err := strconv.Atoi(string(byteVal))
				if err != nil {
					panic(err)
				}
				cards[i] = v
				continue
			}
			switch byteVal {
			case 'A':
				cards[i] = 14
			case 'K':
				cards[i] = 13
			case 'Q':
				cards[i] = 12
			case 'J':
				cards[i] = 11
				if hasJokers {
					cards[i] = 1
				}
			case 'T':
				cards[i] = 10
			}
		}

		hands[i] = Hand{
			cards: cards,
			bid:   b,
			rank:  getHandRank(&cards),
		}
	}

	return hands
}

func getHandRank(h *[]int) int {
	cMap := make(map[int]int)

	wildCard := 0

	for _, b := range *h {
		if b == 1 {
			wildCard++
			continue
		}
		cMap[b]++
	}
	//if wildCard == 5 {
	//	return 0
	//}

	if wildCard > 0 {
		highestDupeIndex := 0
		highestDupe := 0
		for k, v := range cMap {
			if v > highestDupe {
				highestDupeIndex = k
				highestDupe = v
			}
		}
		cMap[highestDupeIndex] += wildCard
	}

	switch len(cMap) {
	case 1:
		return FiveOfAKind
	case 2:
		for c := range cMap {
			if cMap[c] == 4 {
				return FourOfAKind
			}
		}
		return FullHouse
	case 3:
		for k := range cMap {
			if cMap[k] == 3 {
				return ThreeOfAKind
			}
		}
		return TwoPair
	case 4:
		return OnePair
	case 5:
		return HighCard
	}
	panic("should not get here")
}
