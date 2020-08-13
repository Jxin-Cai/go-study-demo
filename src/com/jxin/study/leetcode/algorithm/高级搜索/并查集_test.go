package 高级搜索

import (
	"testing"
)

type UnionFindSet struct {
	count  int
	parent []int
}

func NewUnionFindSet(count int) *UnionFindSet {
	parent := make([]int, count)
	for i := 0; i < count; i++ {
		parent[i] = i
	}
	return &UnionFindSet{count: count, parent: parent}
}

func (s *UnionFindSet) Union(a int, b int) {
	x := s.find(a)
	y := s.find(b)
	if x == y {
		return
	}
	s.parent[x] = y
	s.count--
}

func (s *UnionFindSet) find(a int) int {
	for ; a != s.parent[a]; a = s.parent[a] {
	}
	return a
}

func (s *UnionFindSet) getCount() int {
	return s.count
}

/**
547-朋友圈
*/
func findCircleNum(M [][]int) int {
	if M == nil || len(M) == 0 {
		return 0
	}
	unionFindSet := NewUnionFindSet(len(M))

	for i := 0; i < len(M); i++ {
		for j := 0; j < len(M); j++ {
			if M[i][j] == 0 {
				continue
			}
			unionFindSet.Union(i, j)
		}
	}

	return unionFindSet.getCount()
}
func TestNumIslands(t *testing.T) {
	val := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	numIslands(val)
}

/**
200-岛屿
*/
func numIslands(grid [][]byte) int {
	if grid == nil || len(grid) == 0 {
		return 0
	}
	row, col := len(grid), len(grid[0])
	l := row * col
	unionFindSet := NewUnionFindSet(l + 1)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '0' {
				unionFindSet.Union(i*col+j, l)
				continue
			}
			if i+1 < row && grid[i+1][j] == '1' {
				unionFindSet.Union(i*col+j, (i+1)*col+j)
			}
			if j+1 < col && grid[i][j+1] == '1' {
				unionFindSet.Union(i*col+j, i*col+j+1)
			}
		}
	}
	return unionFindSet.getCount() - 1
}
