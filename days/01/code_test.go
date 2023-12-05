package _1

import "testing"

func Test_getCalibrationTotal(t *testing.T) {
	tests := []struct {
		name                    string
		inputFileName           string
		allowsSpelledOutLetters bool
		want                    int
	}{
		{"Example", "example.txt", false, 142},
		{"Puzzle Input", "input.txt", false, 54605},
		{"Puzzle Input work", "input_work.txt", false, 55477},
		{"Allowing spelled out letters", "spelled_out_letters.txt", true, 495},
		{"Example 2", "example02.txt", true, 281},
		{"Puzzle Input 2", "input.txt", true, 55429},
		{"Puzzle Input 2 work", "input_work.txt", true, 54431},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getCalibrationTotal(tt.inputFileName, tt.allowsSpelledOutLetters); got != tt.want {
				t.Errorf("getCalibrationTotal() = %v, want %v", got, tt.want)
			}
		})
	}
}
