package _1

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func getCalibrationTotal(inputFileName string, allowSpelledOutLetters bool) int {
	content := getFileContent(inputFileName)

	runningTotal := 0
	for _, b := range splitByteArray(content, 10) {
		runningTotal += getLineNumber(b, allowSpelledOutLetters)
	}

	fmt.Printf("Total: %d\n", runningTotal)
	return runningTotal
}

func getLineNumberFromBytes(b []byte) int {
	str := string(b)
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func isDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func isLineBreak(b byte) bool {
	return b == 10
}

func isPotentiallySpelledOutNumber(b byte) bool {
	// o, t, f, s, e, n, O, T, F, S E, N
	// TODO: Find a more idiomatic way to do this
	return b == 111 || b == 116 || b == 102 || b == 115 || b == 101 || b == 110 || b == 79 || b == 84 || b == 70 || b == 83 || b == 69 || b == 78
}

func getSpelledOutNumber(i int, b []byte) (byte, error) {
	// one, two, three, four, five, six, seven, eight, nine, ten
	if i+2 >= len(b) {
		return 0, fmt.Errorf("not a spelled out number")
	}

	str := string(b[i : i+3])
	switch str {
	case "one":
		return 49, nil
	case "two":
		return 50, nil
	case "six":
		return 54, nil
	}

	if i+3 < len(b) {
		str := string(b[i : i+4])
		switch str {
		case "four":
			return 52, nil
		case "five":
			return 53, nil
		case "nine":
			return 57, nil
		}
	}
	if i+4 < len(b) {
		str := string(b[i : i+5])
		switch str {
		case "three":
			return 51, nil
		case "seven":
			return 55, nil
		case "eight":
			return 56, nil
		}
	}

	return 0, fmt.Errorf("not a spelled out number")
}
func getFileContent(inputFileName string) []byte {
	content, err := ioutil.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	return content
}

func splitByteArray(array []byte, separator byte) [][]byte {
	var result [][]byte
	startIndex := 0
	for i := 0; i < len(array); i++ {
		if array[i] == separator {
			result = append(result, array[startIndex:i])
			startIndex = i + 1
		}
	}
	result = append(result, array[startIndex:len(array)])
	return result
}

func getLineNumber(line []byte, allowSpelledOutLetters bool) int {
	currentLineNumberArray := make([]byte, 0)
	for i, b := range line {
		if isDigit(b) {
			if len(currentLineNumberArray) == 2 {
				currentLineNumberArray[1] = b
				continue
			}
			currentLineNumberArray = append(currentLineNumberArray, b, b)
		}
		if allowSpelledOutLetters && isPotentiallySpelledOutNumber(b) {
			if num, err := getSpelledOutNumber(i, line); err == nil {
				if len(currentLineNumberArray) == 2 {
					currentLineNumberArray[1] = num
					continue
				}
				currentLineNumberArray = append(currentLineNumberArray, num, num)
			}
			continue
		}
		//TODO: Don't thing we need this anymore
		if isLineBreak(b) {
			return getLineNumberFromBytes(currentLineNumberArray)
		}
	}
	return getLineNumberFromBytes(currentLineNumberArray)
}
