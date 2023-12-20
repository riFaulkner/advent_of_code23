package _17

import (
	"advent_of_code23/utils"
	"bytes"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	f          int // the f value of the search algo. f = g + h
	g          int // cost to get to this nodea
	h          int // heuristic value, how many moves to get to the final dest
	r          int // row
	c          int // column
	d          int // direction
	con        int // continuous in that direction
	connection *Node
}

func (n Node) key() string {
	return strings.Join([]string{fmt.Sprint(n.c), fmt.Sprint(n.r), fmt.Sprint(n.d), fmt.Sprint(n.con)}, ",")
}

func (n Node) getSuccessors(prob2 bool) []Node {
	dirs := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	successors := make([]Node, 0)

	for i, dir := range dirs {
		newR := n.r + dir[0]
		newC := n.c + dir[1]
		newDir := i
		newCons := 1
		if n.d == newDir {
			newCons = n.con + 1
			if prob2 && newCons > 10 {
				continue
			}
		} else {
			if prob2 && n.con >= 0 && n.con < 4 {
				continue
			}
		}

		if !prob2 && newCons > 3 {
			continue
		}

		// h/t jp -- not sure if this can be updated though to be more better
		isntReverse := ((newDir + 2) % 4) != n.d

		if !isntReverse {
			continue
		}
		successors = append(successors, Node{r: newR, c: newC, h: 0, d: newDir, con: newCons})
	}
	return successors
}

func Problem1(inputFileName string, prob2 bool) int {
	debug := false
	grid := bytes.Split(utils.GetFileContent(inputFileName), []byte("\n"))
	gC := len(grid[0]) - 1
	gR := len(grid) - 1

	// TODO: make this a list of pointers to nodes and or make it a heap and run benchmarks before and after
	open := []Node{{r: 0, c: 0, h: 0, d: -1, con: -1}}

	// TODO: same as above
	closed := make(map[string]Node, 0)

	for len(open) > 0 {
		if debug && len(closed)%10_000 == 0 {
			fmt.Printf("Closed %d nodes\n", len(closed))
		}
		q := open[0]

		in := 0
		for i, n := range open {
			if n.f < q.f || n.f == q.f && n.h < q.h {
				q = n
				in = i
			}
		}
		open = append(open[:in], open[in+1:]...)

		closed[q.key()] = q

		// we're at the final destination
		if q.r == gR && q.c == gC {
			fmt.Printf("End found in val of %d\n", q.g)
			retVal := q.g
			if debug {
				for q.connection != nil {
					grid[q.r][q.c] = '#'
					q = *q.connection
				}

				for _, row := range grid {
					fmt.Println(string(row))
				}
			}

			return retVal
		}

		for _, successor := range q.getSuccessors(prob2) {
			if successor.r < 0 || successor.c < 0 || successor.r > gR || successor.c > gC {
				// successor is out of bounds
				continue
			}

			if _, ok := closed[successor.key()]; ok {
				if closed[successor.key()].g < successor.g {
					print("Found closed item with worse g value")
				}
				continue
			}

			inSearch := slices.IndexFunc(open, func(j Node) bool {
				if successor.r == j.r && successor.c == j.c && successor.d == j.d && successor.con == j.con {
					return true
				}
				return false
			})

			co, err := strconv.Atoi(string(grid[successor.r][successor.c]))
			if err != nil {
				panic(err)
			}
			cost := q.g + co

			if inSearch == -1 || cost < open[inSearch].g {
				if inSearch == -1 {
					successor.g = cost
					successor.h = getDistance(q, gR, gC)
					successor.f = cost + successor.h
					open = append(open, successor)
				} else {
					open[inSearch].g = cost
					open[inSearch].connection = &q
					open[inSearch].f = cost + open[inSearch].h
				}
			}
		}
	}

	return 0
}

func getDistance(a Node, r, c int) int {
	return r - a.r + c - a.c
}
