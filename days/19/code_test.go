package _19

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	type args struct {
		inputFileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example 1", args{"example.txt"}, 19114},
		{"Puzzle 1", args{"puzzle.txt"}, 373120}, // 373120 is too high
		{"Puzzle 1 work", args{"puzzle_work.txt"}, 382440},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName); got != tt.want {
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
		{"Puzzle 1", args{"puzzle.txt"}, 124167549767307},
		{"Puzzle Work", args{"puzzle_work.txt"}, 136394217540123},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem2(tt.args.inputFileName); got != tt.want {
				t.Errorf("Problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}
