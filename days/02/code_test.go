package _2

import "testing"

func Test_getSumOfPossibleGames(t *testing.T) {
	tests := []struct {
		name          string
		inputFileName string
		roundMaxes    Round
		want          int
	}{
		{"Example 1", "example.txt", Round{
			red:   12,
			green: 13,
			blue:  14},
			8},
		{"Puzzle 1", "puzzle_work.txt", Round{
			red:   12,
			green: 13,
			blue:  14},
			2716,
		},
		{"Puzzle 1 Work", "puzzle01_work.txt", Round{
			red:   12,
			green: 13,
			blue:  14},
			2162},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSumOfPossibleGames(tt.inputFileName, tt.roundMaxes); got != tt.want {
				t.Errorf("getSumOfPossibleGames() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getPowerOfMinCubesPossible(t *testing.T) {
	tests := []struct {
		name          string
		inputFileName string
		want          int
	}{
		{"Example 1", "example.txt", 2286},
		{"Puzzle", "puzzle_work.txt", 72227},
		{"Puzzle 1 Work", "puzzle01_work.txt", 72513},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPowerOfMinCubesPossible(tt.inputFileName); got != tt.want {
				t.Errorf("getPowerOfMinCubesPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
