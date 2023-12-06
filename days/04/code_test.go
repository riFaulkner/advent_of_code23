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
		{"Puzzle Input", args{"input_work.txt"}, 21821},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTotalScratchCardPoints(tt.args.inputFileName); got != tt.want {
				t.Errorf("GetTotalScratchCardPoints() = %v, want %v", got, tt.want)
			}
		})
	}
}
