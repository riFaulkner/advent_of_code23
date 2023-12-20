package _4

import "testing"

func TestGetTotalScratchCardPoints(t *testing.T) {
	type args struct {
		inputFileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{"example.txt"}, 13},
		{"Puzzle Input", args{"puzzle_work.txt"}, 20667},
		{"Puzzle Input Work", args{"input_work.txt"}, 21821},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTotalScratchCardPoints(tt.args.inputFileName); got != tt.want {
				t.Errorf("GetTotalScratchCardPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTotalScratchCards(t *testing.T) {
	type args struct {
		inputFileName string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{"example.txt"}, 30},
		{"Puzzle Input", args{"puzzle_work.txt"}, 5833065},
		{"Puzzle Input Work", args{"input_work.txt"}, 5539496},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTotalScratchCards(tt.args.inputFileName); got != tt.want {
				t.Errorf("GetTotalScratchCards() = %v, want %v", got, tt.want)
			}
		})
	}
}
