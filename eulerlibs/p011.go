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
