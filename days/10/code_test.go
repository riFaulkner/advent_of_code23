package _0

import "testing"

func TestProblem1(t *testing.T) {
	type args struct {
		inputFileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{"example01.txt"}, 4},
		{"Example 2", args{"example02.txt"}, 8},
		{"Puzzle 1", args{"puzzle_work.txt"}, 6757},
		{"Puzzle 1 work", args{"puzzle_work.txt"}, 6867},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}
