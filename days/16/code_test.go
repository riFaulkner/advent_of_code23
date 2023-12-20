package _16

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
		{"Example", args{"example.txt"}, 46},
		{"Puzzle", args{"puzzle_work.txt"}, 7608},
		{"Puzzle Work", args{"puzzle_work.txt"}, 7632},
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
		{"Example", args{"example.txt"}, 51},
		{"Puzzle", args{"puzzle.txt"}, 8221},
		{"Puzzle Work", args{"puzzle_work.txt"}, 8023},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem2(tt.args.inputFileName); got != tt.want {
				t.Errorf("Problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}
