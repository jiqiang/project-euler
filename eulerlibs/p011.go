package eulerlibs

// HorizontalCheck retuns product of 4 adjcent numbers horizontally.
func HorizontalCheck(grid [20][20]int64) int64 {
	var p, max int64
	max = 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j+3 < len(grid[i]); j++ {
			if grid[i][j] == 0 || grid[i][j+1] == 0 || grid[i][j+2] == 0 || grid[i][j+3] == 0 {
				continue
			}
			p = grid[i][j] * grid[i][j+1] * grid[i][j+2] * grid[i][j+3]
			if max < p {
				max = p
			}
		}
	}
	return max
}

// VerticalCheck retuns product of 4 adjcent numbers vertically.
func VerticalCheck(grid [20][20]int64) int64 {
	var p, max int64
	max = 0
	for i := 0; i+3 < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 0 || grid[i+1][j] == 0 || grid[i+2][j] == 0 || grid[i+3][j] == 0 {
				continue
			}
			p = grid[i][j] * grid[i+1][j] * grid[i+2][j] * grid[i+3][j]
			if max < p {
				max = p
			}
		}
	}
	return max
}
