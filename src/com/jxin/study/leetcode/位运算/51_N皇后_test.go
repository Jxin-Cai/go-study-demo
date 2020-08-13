package 位运算

import "testing"

func TestQueen(t *testing.T) {
	solveNQueens(4)
}

var ret [][]string

func solveNQueens(n int) [][]string {
	ret = nil
	dfs(0, n, 0, 0, 0, []int{})
	return ret
}

func dfs(offset int, n int, col int, left int, right int, count []int) {
	if offset == n {
		return
	}
	size := (1 << n) - 1

	validSite := size & (^(col | left | right))

	for validSite != 0 {
		site := validSite & (-validSite)
		count = append(count, site)
		dfs(offset+1, n, col|site, (left|site)<<1, (right|site)>>1, count)
		if offset+1 == n {
			addRet(count)
		}
		count = count[0:offset]
		validSite &= validSite - 1
	}
}

func addRet(count []int) {
	n := len(count)
	oneRet := make([]string, n)
	for i := 0; i < n; i++ {
		row := ""
		col := countSize(count[i])
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

func countSize(i int) int {
	ret := 0
	if i == 1 {
		return ret
	}
	for ; i > 1; i >>= 1 {
		ret++
	}
	return ret
}
