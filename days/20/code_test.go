package _20

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
		{"example 1", args{"example.txt"}, 32000000},
		{"example 2", args{"example2.txt"}, 11687500},
		{"Puzzle 1 work", args{"puzzle_work.txt"}, 800830848},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Problem1(tt.args.inputFileName); got != tt.want {
				t.Errorf("Problem1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProcessPulse(t *testing.T) {
	p := 0
	for i := 0; i < 10; i++ {
		p = (p + 1) % 2
		t.Log(i, p, (p+1)%2)
	}
}

//button -low-> broadcaster
//broadcaster -low-> a
//broadcaster -low-> b
//broadcaster -low-> c
//a -high-> b
//b -high-> c
//c -high-> inv
//inv -low-> a
//a -low-> b
//b -low-> c
//c -low-> inv
//inv -high-> a

//func TestProblem2(t *testing.T) {
//	type args struct {
//		inputFileName string
//	}
//	tests := []struct {
//		name string
//		args args
//		want int
//	}{
//		//{"Puzzle work", args{"puzzle_work.txt"}, 0},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := Problem2(tt.args.inputFileName); got != tt.want {
//				t.Errorf("Problem2() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}
