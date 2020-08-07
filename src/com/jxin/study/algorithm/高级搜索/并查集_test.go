package 高级搜索

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
