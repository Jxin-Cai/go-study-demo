package 染色算法

import (
	"testing"
)

func TestNumIslands(t *testing.T) {
	val := [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	numIslands(val)
}

func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	ret := 0
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '0' {
				continue
			}
			ret++
			bfs(i, j, grid)
		}
	}
	return ret
}

func dfs(i int, j int, grid [][]byte) {
	row, col := len(grid), len(grid[0])
	if i >= row || i < 0 {
		return
	}
	if j >= col || j < 0 {
		return
	}
	if grid[i][j] == '0' {
		return
	}
	grid[i][j] = '0'
	dfs(i+1, j, grid)
	dfs(i-1, j, grid)
	dfs(i, j+1, grid)
	dfs(i, j-1, grid)
}

func bfs(i int, j int, grid [][]byte) {
	row, col := len(grid), len(grid[0])
	queue := make([]int, 0)
	queue = append(queue, i, j)
	for len(queue) != 0 {
		l := len(queue)
		for z := 0; z < l/2; z++ {
			x, y := queue[0], queue[1]
			queue = queue[2:]
			if grid[x][y] == '0' {
				continue
			}
			grid[x][y] = '0'
			if x+1 < row {
				queue = append(queue, x+1, y)
			}
			if x-1 >= 0 {
				queue = append(queue, x-1, y)
			}
			if y+1 < col {
				queue = append(queue, x, y+1)
			}
			if y-1 >= 0 {
				queue = append(queue, x, y-1)
			}
		}
	}
}
