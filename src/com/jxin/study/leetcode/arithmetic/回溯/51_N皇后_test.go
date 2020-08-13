package 回溯

import "testing"

func TestQueen(t *testing.T) {
	solveNQueens(4)
}

var col, left, right []int
var ret [][]string

func solveNQueens(n int) [][]string {
	if n == 0 {
		return nil
	}
	col = make([]int, n)
	left = make([]int, n*2-1)
	right = make([]int, n*2-1)
	queen := make([]int, n)
	ret = nil
	dfs(n, 0, queen)
	return ret
}

func dfs(n int, offset int, queen []int) {
	if offset == n {
		return
	}
	for i := 0; i < n; i++ {
		if !addQueen(offset, i, n, queen) {
			continue
		}
		if offset+1 == n {
			addRet(queen, n)
		}
		dfs(n, offset+1, queen)
		reset(offset, i, n, queen)
	}
}

func addQueen(x int, y int, n int, queen []int) bool {
	if col[y] == 1 {
		return false
	}
	if left[n-1-y+x] == 1 {
		return false
	}
	if right[x+y] == 1 {
		return false
	}
	col[y] = 1
	left[n-1-y+x] = 1
	right[x+y] = 1
	queen[x] = y
	return true
}

func reset(x int, y int, n int, queen []int) {
	col[y] = 0
	left[n-1-y+x] = 0
	right[x+y] = 0
	queen = queen[0:x]
}

func addRet(queen []int, n int) {
	oneRet := make([]string, n)
	for i := 0; i < n; i++ {
		row := ""
		col := queen[i]
		for j := 0; j < n; j++ {
			if j == col {
				row += "Q"
				continue
			}
			row += "."
		}
		oneRet[i] = row
	}
	ret = append(ret, oneRet)
}
