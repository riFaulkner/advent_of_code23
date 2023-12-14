package _12

import (
	"advent_of_code23/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var moves = []byte{'#', '.'}

func Problem1(inputFileName string, expansionMultiplier int) int {
	lines := strings.Split(utils.GetFileAsString(inputFileName), "\n")
	c := make(chan int)

	sum := 0
	for _, line := range lines {
		go calculateDifferentArrangements(line, expansionMultiplier, c)
	}
	fmt.Printf("Progress Out of %d\n", len(lines))
	for i := 0; i < len(lines); i++ {
		fmt.Printf("|")
		sum += <-c
	}
	fmt.Printf("\n DONE \n")
	return sum
}

func calculateDifferentArrangements(line string, expansionMultiplier int, c chan int) {
	segs := strings.Split(line, " ")

	l := make([]string, expansionMultiplier)
	m := make([]string, expansionMultiplier)

	for i := 0; i < expansionMultiplier; i++ {
		l[i] = segs[0]
		m[i] = segs[1]
	}
	expandedLine := strings.Join(l, "?")
	expandedMap := strings.Join(m, ",")

	keyB := regexp.MustCompile(`(\d)+`).FindAllString(expandedMap, -1)
	keyList := make([]int, len(keyB))
	for i := range keyB {
		intVal, err := strconv.Atoi(keyB[i])
		if err != nil {
			panic(err)
		}
		keyList[i] = intVal
	}

	dp := make(map[string]int)

	c <- getMoveCountDP([]byte(expandedLine+"."), keyList, 0, 0, 0, &dp)
}

func getMoveCountDP(line []byte, key []int, bi, ki, curr int, dp *map[string]int) int {
	dpKey := strings.Join([]string{string(rune(bi)), string(rune(ki)), string(rune(curr))}, "-")
	if val, ok := (*dp)[dpKey]; ok {
		// we have the value cached
		return val
	}
	if bi == len(line) { // we're at the end of the line, check conditions
		if ki == len(key) && curr == 0 { // We've already completed all blocks and we don't have a current run we're working on
			return 1
		} else if ki == len(key)-1 && curr == key[ki] { // We're at the last block and our current run is the same size as the last block
			return 1
		}
		return 0
	}

	moveCount := 0
	for _, c := range moves {
		if line[bi] == c || line[bi] == '?' {
			switch c {
			case '.':
				// Just move the byte
				if curr == 0 {
					moveCount += getMoveCountDP(line, key, bi+1, ki, curr, dp)
				}
				if curr > 0 && ki < len(key) && curr == key[ki] {
					// we need to close out the last block.
					moveCount += getMoveCountDP(line, key, bi+1, ki+1, 0, dp)
				}
			case '#':
				// Potential optimization is to look now and see if curr +1 > key[ki] if it is noop/return
				moveCount += getMoveCountDP(line, key, bi+1, ki, curr+1, dp)
			}
		}
	}

	(*dp)[dpKey] = moveCount

	return moveCount

}
