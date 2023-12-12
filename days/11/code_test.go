package _11

import "testing"

func TestProblem1(t *testing.T) {
	type args struct {
		inputFileName       string
		expansionMultiplier int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example ", args{"example.txt", 1}, 374},
		{"Puzzle 1", args{"puzzle_work.txt", 1}, 10077850},
		{"Example high multiplier ", args{"example.txt", 9}, 1030},
		{"Example high high multiplier ", args{"example.txt", 99}, 8410},
		{"Puzzle 1 high high higher multiplier", args{"puzzle_work.txt", 999_999}, 504715068438},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName, tt.args.expansionMultiplier); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}
