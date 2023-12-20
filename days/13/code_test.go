package _13

import "testing"

func TestProblem1(t *testing.T) {
	type args struct {
		inputFileName string
		failures      int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example ", args{"example.txt", 0}, 405},
		{"Puzzle 1", args{"puzzle_work.txt", 0}, 33520},
		{"Puzzle 1 work", args{"puzzle_work.txt", 0}, 36448},
		{"Example smudges", args{"example.txt", 1}, 400},
		{"Puzzle 2", args{"puzzle_work.txt", 1}, 34824},
		{"Puzzle 2 work", args{"puzzle_work.txt", 1}, 35799},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName, tt.args.failures); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}
