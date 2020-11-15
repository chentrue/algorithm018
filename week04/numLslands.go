func numIslands(grid [][]byte) int {
	count := 0
	row, col := len(grid), len(grid[0])
	for i:= 0; i < row; i++{
		for j:=0; j < col; j++{
			if grid[i][j] == '1' {

				count ++
				worker(grid, i, j)
			}
		}
	}
	return count
}

func worker(grid [][]byte, i,j int ) {
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == '0'{
		return
	}
	grid[i][j] = '0'
	worker(grid, i+1, j)
	worker(grid, i, j+1)
	worker(grid, i-1, j)
	worker(grid, i, j-1)
}
