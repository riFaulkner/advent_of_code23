package _19

import "testing"

func TestProblem1(t *testing.T) {
	type args struct {
		inputFileName string
		part2         bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"example.txt", false}, 19114},
		{"Puzzle 1 work", args{"puzzle_work.txt", false}, 382440},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName, tt.args.part2); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProblem2(t *testing.T) {
	type args struct {
		inputFileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"example.txt"}, 167409079868000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem2(tt.args.inputFileName); got != tt.want {
				t.Errorf("Problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}
