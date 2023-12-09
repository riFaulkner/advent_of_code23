package _5

import (
	"advent_of_code23/utils"
	"cmp"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"time"
)

// Started this out with brute force, it kept failing after 2:20 minutes
// After making the seed seiralization non-struct based it finished in

// First attempt with channels is now causing a lock on lower numbers of seeds but is processing at ~6s per billion
// "Finished after 20 min 26 seconds

// Added a local min to control what gets sent through the channel
// took it down to 1:03 min

// Adding in additional batching brought it down to 30 seconds

type Almanac struct {
	sourceDestinationMaps map[string]SourceDestinationMapping
}

type SourceDestinationMapping struct {
	source          string
	destination     *string
	adjustmentsList []Adjustment
}

type Adjustment struct {
	sourceStart      int
	destinationStart int
	offset           int
}

func GetClosestSeedPlaningLocation(inputFileName string, seedsAsRange bool) int {
	content := utils.GetFileContent(inputFileName)
	sections := regexp.MustCompile(`\n\s*\n`).Split(string(content), -1)

	almanac := serializeAlmanac(sections[1:])

	return findClosestSeedLocation(almanac, sections[0], seedsAsRange)
}

func seedsToChannels(almanac Almanac, content string, seedsAsRange bool) []chan int {
	channels := make([]chan int, 0)
	start := time.Now()

	numbers := regexp.MustCompile(`\d+`).FindAllString(content, -1)
	// split by empty lines
	for i, number := range numbers {
		if seedsAsRange && i%2 == 1 {
			continue
		}
		//c := make(chan int)
		// turn number into int
		intNum, err := strconv.Atoi(number)
		if err != nil {
			panic(err)
		}
		offset := 0
		numChunks := 1
		if seedsAsRange {
			offset, err = strconv.Atoi(numbers[i+1])
			if err != nil {
				panic(err)
			}
			numChunks = 10
		}
		chunkSize := offset / numChunks
		for ch := 0; ch < numChunks; ch++ {
			c := make(chan int)

			startVal := intNum + (ch * chunkSize)

			go processChunks(c, almanac, startVal, chunkSize)
			channels = append(channels, c)
		}

		//channels = append(channels, c)
	}
	fmt.Printf("%d Channels created in %s\n", len(channels), time.Since(start))

	return channels
}

func processChunks(c chan int, a Almanac, start, offset int) {
	localMin := 9999999999
	seeds := 0
	startTime := time.Now()
	seedS := "seed"
	locationS := "location"

	for j := 0; j <= offset; j++ {
		seeds++
		currentSeedDistance := a.followToDestination(start+j, &seedS, &locationS)
		if currentSeedDistance < localMin {
			localMin = currentSeedDistance
		}
	}
	end := time.Since(startTime)
	fmt.Printf("Finished processing %d seeds on goroutine in %v: %dms avg\n", seeds, end, int64(seeds)/end.Nanoseconds())
	c <- localMin

}

func fanInChannels(channels []chan int) (chan int, int) {
	c := make(chan int)
	for _, channel := range channels {
		go func(channel chan int) {
			for {
				c <- <-channel
			}
		}(channel)
	}
	return c, len(channels)
}

func serializeAlmanac(content []string) Almanac {
	sourceDestinationMappings := make(map[string]SourceDestinationMapping)
	for _, line := range content {
		s := regexp.MustCompile(`([ map]+):`).Split(line, -1)
		sourceAndDestination := regexp.MustCompile(`([-])+`).Split(s[0], -1)

		mappingStrings := regexp.MustCompile(`\b[^\n]*`).FindAllString(s[1], -1)
		adjustments := make([]Adjustment, 0)
		for _, mappingString := range mappingStrings {
			numbers := regexp.MustCompile(`\d+`).FindAllString(mappingString, -1)

			adjustments = append(adjustments, Adjustment{
				destinationStart: utils.GetNumberFromBytes([]byte(numbers[0])),
				sourceStart:      utils.GetNumberFromBytes([]byte(numbers[1])),
				offset:           utils.GetNumberFromBytes([]byte(numbers[2])),
			})
		}

		slices.SortFunc(adjustments, func(a, b Adjustment) int {
			return cmp.Compare(a.sourceStart, b.sourceStart)
		})

		mapping := SourceDestinationMapping{
			source:          sourceAndDestination[0],
			destination:     &sourceAndDestination[2],
			adjustmentsList: adjustments,
		}

		sourceDestinationMappings[sourceAndDestination[0]] = mapping
	}

	return Almanac{
		sourceDestinationMaps: sourceDestinationMappings,
	}
}

func (a Almanac) findClosestSeedDistance(c chan int, numC int, startTime time.Time) int {
	numSeeds := 0

	// max int
	closestDistance := 9999999999

	for i := 0; i < numC; i++ {
		s := <-c
		numSeeds++
		if s < closestDistance {
			closestDistance = s
		}

	}

	fmt.Printf("No more inputs found Processed %d seeds in %v \n", numSeeds, time.Since(startTime))
	return closestDistance
}

func findClosestSeedLocation(almanac Almanac, content string, seedsAsRange bool) int {
	start := time.Now()
	c, numC := fanInChannels(seedsToChannels(almanac, content, seedsAsRange))
	fmt.Printf("All Channels created in %s\n", time.Since(start))
	return almanac.findClosestSeedDistance(c, numC, start)
}

func (a Almanac) followToDestination(number int, source, finalDest *string) int {
	nextSource := source
	num := number
	for {
		mapping := a.sourceDestinationMaps[*nextSource]

		destinationValue := calculateDestinationValue(num, mapping.adjustmentsList)

		if mapping.destination == nil || *mapping.destination == *finalDest {
			return destinationValue
		}

		num = destinationValue
		nextSource = mapping.destination
	}
}

func calculateDestinationValue(value int, adjustments []Adjustment) int {
	for _, adjustment := range adjustments {
		if value-adjustment.sourceStart < 0 {
			return value
		}
		if (adjustment.sourceStart+adjustment.offset)-value > 0 {
			// We have a match, the value is in our range
			return adjustment.destinationStart + (value - adjustment.sourceStart)
		}
	}
	return value
}
