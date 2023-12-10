package _7

import (
	"testing"
)

func TestProblem1(t *testing.T) {
	type args struct {
		inputFile string
		hasJokers bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example 01", args{"example.txt", false}, 6440},
		{"Puzzle 01", args{"puzzle_work.txt", false}, 246795406},
		{"Example 02", args{"example.txt", true}, 5905},
		{"Puzzle 02", args{"puzzle_work.txt", true}, 249356515},
	}
	// 246795406 too low...
	// 248298903
	// 248711591
	// 248856588
	// 248861756
	// 248714781
	// 248640674
	// 248876000
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFile, tt.args.hasJokers); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHandRank(t *testing.T) {
	type args struct {
		h *[]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// Five of a kind, where all five cards have the same label: AAAAA
		// Four of a kind, where four cards have the same label and one card has a different label: AA8AA
		// Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
		// Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
		// Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
		// One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
		{"High Card", args{&[]int{2, 3, 4, 5, 6}}, HighCard},
		{"HC wild to a Pair", args{&[]int{1, 3, 4, 5, 6}}, OnePair},
		{"Pair", args{&[]int{3, 4, 5, 6, 3}}, OnePair},
		{"Pair wild to three of a kind", args{&[]int{1, 3, 3, 5, 6}}, ThreeOfAKind},
		{"All Wilds", args{&[]int{1, 1, 1, 1, 1}}, FiveOfAKind},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHandRank(tt.args.h); got != tt.want {
				t.Errorf("getHandRank() = %v, want %v", got, tt.want)
			}
		})
	}
}
