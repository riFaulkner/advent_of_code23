package _15

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
		{"Example 1", args{"example.txt"}, 1320},
		{"Puzzle", args{"puzzle_work.txt"}, 506869},
		{"Puzzle Work", args{"puzzle_work.txt"}, 501680},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hashString(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{[]byte("HASH")}, 52},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hashString(tt.args.b); got != tt.want {
				t.Errorf("hashString() = %v, want %v", got, tt.want)
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
		{"Example 1", args{"example.txt"}, 145},
		{"Puzzle 2", args{"puzzle_work.txt"}, 271384},
		{"Puzzle 2 Work", args{"puzzle_work.txt"}, 241094},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem2(tt.args.inputFileName); got != tt.want {
				t.Errorf("Problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}
