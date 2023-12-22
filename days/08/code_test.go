package _8

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
		{"example", args{"example01.txt"}, 2},
		{"example 2", args{"example02.txt"}, 6},
		{"Puzzle", args{"puzzle_work.txt"}, 13771},
		{"Puzzle work", args{"puzzle_work.txt"}, 20221},
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
		{"example", args{"example03.txt"}, 6},
		{"Puzzle", args{"puzzle.txt"}, 13129439557681},
		{"Puzzle Work", args{"puzzle_work.txt"}, 14616363770447},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem2(tt.args.inputFileName); got != tt.want {
				t.Errorf("Problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findLowestCommonMultiple(t *testing.T) {
	type args struct {
		vals []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"easy example", args{[]int{2, 3}}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findLowestCommonMultiple(tt.args.vals); got != tt.want {
				t.Errorf("findLowestCommonMultiple() = %v, want %v", got, tt.want)
			}
		})
	}
}
