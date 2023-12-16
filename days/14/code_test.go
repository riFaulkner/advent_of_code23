package _14

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
		{"Example 1", args{"example.txt"}, 136},
		{"Puzzle 1", args{"puzzle_work.txt"}, 110407},
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
		{"Example 1", args{"example.txt"}, 64},
		{"Puzzle 2", args{"puzzle_work.txt"}, 87273},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem2(tt.args.inputFileName); got != tt.want {
				t.Errorf("Problem2() = %v, want %v", got, tt.want)
			}
		})
	}
}

//After 1 cycle:
//.....#....
//....#...O#
//...OO##...
//.OO#......
//.....OOO#.
//.O#...O#.#
//....O#....
//......OOOO
//#...O###..
//#..OO#....
//
//After 2 cycles:
//.....#....
//....#...O#
//.....##...
//..O#......
//.....OOO#.
//.O#...O#.#
//....O#...O
//.......OOO
//#..OO###..
//#.OOO#...O
//
//After 3 cycles:
//.....#....
//....#...O#
//.....##...
//..O#......
//.....OOO#.
//.O#...O#.#
//....O#...O
//.......OOO
//#...O###.O
//#.OOO#...O
