package utils

import "strconv"

func GetNumberFromBytes(b []byte) int {
	str := string(b)
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func IsComma(b byte) bool {
	return b == 44
}

func IsDigit(b byte) bool {
	return b >= 48 && b <= 57
}

func IsLetter(b byte) bool {
	return b >= 97 && b <= 122
}

func IsLineBreak(b byte) bool {
	return b == 10
}

func IsSemiColon(b byte) bool {
	return b == 59
}
