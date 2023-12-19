package _17

import "testing"

func TestProblem1(t *testing.T) {
	type args struct {
		inputFileName string
		prob2         bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 1", args{"example.txt", false}, 102},
		{"Puzzle 1 Work", args{"puzzle_work.txt", false}, 843}, // 844 is to high
		{"Example 2", args{"example.txt", true}, 94},
		{"Puzzle 2 Work", args{"puzzle_work.txt", true}, 1017}, // 844 is to high
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName, tt.args.prob2); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// 2##34####1323	2>>34^>>>1323
// 32####35##623    32v>>>35v5623
// 325524565##54    32552456v>>54
// 3446585845###    3446585845v52
// 454665786753#    4546657867v>6
// 14385987984##    14385987984v4
// 44578769877#6    44578769877v6
// 36378779796##    36378779796v>
// 465496798688#    465496798688v
// 456467998645#    456467998645v
// 12246868655##    12246868655<v
// 25465488877#5    25465488877v5
// 43226746555##    43226746555v>
