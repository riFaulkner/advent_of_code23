package _17

import "testing"

func TestProblem1(t *testing.T) {
	type args struct {
		inputFileName string
		prob2         bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"example.txt", false}, 102},
		{"Puzzle 1 Work", args{"puzzle_work.txt", false}, 843}, // 844 is to high
		{"Example 2", args{"example.txt", true}, 94},
		{"Puzzle 2 Work", args{"puzzle_work.txt", true}, 1017}, // 844 is to high
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName, tt.args.prob2); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}
