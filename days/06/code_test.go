package _6

import "testing"

func TestGetPowerOfWaysToWin(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 01", args{"example.txt"}, 288},
		{"Puzzle", args{"puzzle.txt"}, 2612736},
		{"Puzzle Work", args{"puzzle01_work.txt"}, 1624896},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPowerOfWaysToWin(tt.args.inputFile); got != tt.want {
				t.Errorf("GetPowerOfWaysToWin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProblemTwo(t *testing.T) {
	type args struct {
		inputFile string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 01", args{"example.txt"}, 71503},
		{"Puzzle", args{"puzzle.txt"}, 29891250},
		{"Puzzle Work", args{"puzzle01_work.txt"}, 32583852},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProblemTwo(tt.args.inputFile); got != tt.want {
				t.Errorf("ProblemTwo() = %v, want %v", got, tt.want)
			}
		})
	}
}
