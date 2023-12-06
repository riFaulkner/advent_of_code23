package _3

import "testing"

func TestGetSumOfSchematicParts(t *testing.T) {
	type args struct {
		inputFileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{"example.txt"}, 4361},
		{"Puzzle Input 1 Work", args{"puzzle01_work.txt"}, 528799},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSumOfSchematicParts(tt.args.inputFileName); got != tt.want {
				t.Errorf("GetSumOfSchematicParts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSumOfGearRatios(t *testing.T) {
	type args struct {
		inputFileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 2", args{"example.txt"}, 467835},
		{"Puzzle Input 2 Work", args{"puzzle01_work.txt"}, 84907174},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSumOfGearRatios(tt.args.inputFileName); got != tt.want {
				t.Errorf("GetSumOfGearRatios() = %v, want %v", got, tt.want)
			}
		})
	}
}
