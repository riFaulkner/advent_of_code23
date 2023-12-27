package utils

var Directions = [][]int{
	{0, 1},  // right
	{1, 0},  // down
	{0, -1}, // left
	{-1, 0}, // up
}

func IsInbounds(grid *[][]byte, r, c int) bool {
	return c >= 0 && c < len((*grid)[0]) && r >= 0 && r < len(*grid)
}
