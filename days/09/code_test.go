package _9

import "testing"

func TestProblem1(t *testing.T) {
	type args struct {
		inputFileName string
		rev           bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 01", args{"example.txt", false}, 114},
		{"Puzzle 01 work ", args{"puzzle_work.txt", false}, 1708206096},

		{"Example 01", args{"example.txt", true}, 2},
		{"Puzzle 01 work ", args{"puzzle_work.txt", true}, 1050},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName, tt.args.rev); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}
