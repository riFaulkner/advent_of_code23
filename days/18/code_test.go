package _18

import (
	orderedmap "github.com/wk8/go-ordered-map"
	"testing"
)

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
		{"Example", args{"example.txt", false}, 62},
		{"Puzzle", args{"puzzle.txt", false}, 52231},
		{"Puzzle Work", args{"puzzle_work.txt", false}, 74074},
		{"Example 2 ", args{"example.txt", true}, 952408144115},
		{"Puzzle 2", args{"puzzle.txt", true}, 57196493937398},
		{"Puzzle 2 Work", args{"puzzle_work.txt", true}, 112074045986829},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName, tt.args.part2); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcAreaFromCoorSet(t *testing.T) {
	type args struct {
		coorSet func() *orderedmap.OrderedMap
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "example", args: args{func() *orderedmap.OrderedMap {
			om := orderedmap.New()
			// example from online. answer is 45.5 (45 as an int)
			om.Set("4, 10", []int{4, 10})
			om.Set("9, 7", []int{9, 7})
			om.Set("11, 2", []int{11, 2})
			om.Set("2, 2", []int{2, 2})

			return om
		},
		}, want: 45,
		},
		{name: "example2", args: args{func() *orderedmap.OrderedMap {
			om := orderedmap.New()

			om.Set("1, 10", []int{1, 10})
			om.Set("7, 10", []int{7, 10})
			om.Set("7, 5", []int{7, 5})
			om.Set("5, 5", []int{5, 5})
			om.Set("5, 3", []int{5, 3})
			om.Set("7, 3", []int{7, 3})
			om.Set("7, 1", []int{7, 1})
			om.Set("2, 1", []int{2, 1})
			om.Set("2, 3", []int{2, 3})
			om.Set("1, 3", []int{1, 3})

			om.Set("1, 5", []int{1, 5})

			om.Set("3, 5", []int{3, 5})
			om.Set("3, 8", []int{3, 8})
			om.Set("1, 8", []int{1, 8})
			om.Set("1, 10", []int{1, 10})

			return om
		},
		}, want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcAreaFromCoorSet(tt.args.coorSet()); got != tt.want {
				t.Errorf("calcAreaFromCoorSet() = %v, want %v", got, tt.want)
			}
		})
	}
}
