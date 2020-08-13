package 位运算

import "testing"

func TestQueen(t *testing.T) {
	solveNQueens(4)
}

var ret int

func solveNQueens(n int) int {
	ret = 0
	dfs(0, n, 0, 0, 0)
	return ret
}

func dfs(offset int, n int, col int, left int, right int) {
	if offset == n {
		return
	}
	size := (1 << n) - 1

	validSite := size & (^(col | left | right))

	for validSite != 0 {
		site := validSite & (-validSite)
		dfs(offset+1, n, col|site, (left|site)<<1, (right|site)>>1)
		if offset+1 == n {
			ret++
		}
		validSite &= validSite - 1
	}
}
