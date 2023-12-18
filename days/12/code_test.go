package _12

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
		{"Example ", args{"example.txt", 1}, 21},
		{"Puzzle 1", args{"puzzle.txt", 1}, 8419},
		{"Puzzle 1 Work", args{"puzzle_work.txt", 1}, 7633},
		{"Example 2", args{"example.txt", 5}, 525152},
		{"Puzzle 2", args{"puzzle.txt", 5}, 160500973317706},
		{"Puzzle 2 Work", args{"puzzle_work.txt", 5}, 23903579139437},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName, tt.args.expansionMultiplier); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCalculateDifferentArrangements(t *testing.T) {
	type args struct {
		line               string
		exansionMultiplier int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"???.### 1,1,3", 1}, 1},
		{"Example 2", args{".??..??...?##. 1,1,3", 1}, 4},
		{"Example 3", args{"?#?#?#?#?#?#?#? 1,3,1,6", 1}, 1},
		{"Example 4", args{"????.#...#... 4,1,1", 1}, 1},
		{"Example 5", args{"????.######..#####. 1,6,5", 1}, 4},
		{"Example 6", args{"?###???????? 3,2,1", 1}, 10},
		{"Example 1 pt2", args{"???.### 1,1,3", 5}, 1},

		{"Example 2", args{".??..??...?##. 1,1,3", 5}, 16384},
		{"Example 3", args{"?#?#?#?#?#?#?#? 1,3,1,6", 5}, 1},
		{"Example 4", args{"????.#...#... 4,1,1", 5}, 16},
		{"Example 5", args{"????.######..#####. 1,6,5", 5}, 2500},
		{"Example 6", args{"?###???????? 3,2,1", 5}, 506250},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := make(chan int, 1)
			calculateDifferentArrangements(tt.args.line, tt.args.exansionMultiplier, c)

			if got := <-c; got != tt.want {
				t.Errorf("calculateDifferentArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
