package 贪心

import (
	"sort"
	"testing"
)

func TestServic(t *testing.T) {
	findContentChildren([]int{1, 2, 3}, []int{1, 1})

}

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	ret := 0
	for _, v := range g {
		for i := 0; i < len(s); i++ {
			if s[i] < v {
				continue
			}
			ret++
			s[i] = 0
			break
		}
	}
	return ret
}
func findContentChildren1(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	i := 0
	j := 0
	for i < len(g) && j < len(s) {
		if g[i] <= s[j] {
			i++
		}
		j++
	}
	return i
}
