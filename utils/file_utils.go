package utils

import "io/ioutil"

func GetFileContent(inputFileName string) []byte {
	content, err := ioutil.ReadFile(inputFileName)
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
