package utils

import (
	"os"
)

func GetFileAsString(inputFileName string) string {
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	return string(content)
}

func GetFileContent(inputFileName string) []byte {
	content, err := os.ReadFile(inputFileName)
	if err != nil {
		panic(err)
	}
	return content
}

func SplitByteArrayByLine(array []byte) [][]byte {
	return splitByteArray(array, '\n')
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
