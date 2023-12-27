package _21

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestProblem1(t *testing.T) {
	type args struct {
		inputFileName string
		numSteps      int
		part2         bool
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Example", args{"example.txt", 6, false}, 16},
		{"Example 2", args{"example.txt", 10, true}, 50},
		//In exactly 10 steps, he can reach any of 50 garden plots.
		//In exactly 50 steps, he can reach 1594 garden plots.
		//In exactly 100 steps, he can reach 6536 garden plots.
		//In exactly 500 steps, he can reach 167004 garden plots.
		//In exactly 1000 steps, he can reach 668697 garden plots.
		//In exactly 5000 steps, he can reach 16733044 garden plots.
		//{"Example 3", args{"example.txt", 50, true}, 1594},
		{"Puzzle work", args{"puzzle_work.txt", 64, false}, 3795},
		//{"Puzzle work 2", args{"puzzle_work.txt", 26501365, false}, 3795},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName, tt.args.numSteps, tt.args.part2); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

// Proves that fmt.Sprintf is the best
func BenchmarkStringSerialize(b *testing.B) {
	b.Run("strings.Join", func(b *testing.B) {
		benchStringSerialize(1, b)
	})
	b.Run("fmt.Sprintf", func(b *testing.B) {
		benchStringSerialize(2, b)
	})
	b.Run("bytes.Buffer", func(b *testing.B) {
		benchStringSerialize(3, b)
	})

}

func benchStringSerialize(m int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		for i := 0; i < 10_000_000; i++ {
			switch m {
			case 1:
				// strings join
				strings.Join([]string{fmt.Sprint(i), fmt.Sprint(i + 1), fmt.Sprint(i - 1)}, "-")
			case 2:
				// fmt.Sprintf
				fmt.Sprintf("%d-%d-%d", i, i+1, i-1)
			case 3:
				// bytes.Buffer
				var buffer bytes.Buffer
				buffer.WriteString(fmt.Sprint(i))
				buffer.WriteString("-")
				buffer.WriteString(fmt.Sprint(i + 1))
				buffer.WriteString("-")
				buffer.WriteString(fmt.Sprint(i - 1))
				buffer.String()
			}
		}
	}
}
