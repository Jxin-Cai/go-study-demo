package bfs

func ladderLength(beginWord string, endWord string, wordList []string) int {
	if len(beginWord) < 1 || len(beginWord) != len(endWord) {
		return 0
	}

	wdict := initWdict(wordList)

	boolArr := make([]bool, len(wordList))
	count := 1

	wordList = append(wordList, beginWord)
	q := []int{len(wordList) - 1}
	for len(q) > 0 {
		var nextQ []int
		count++
		for _, qid := range q {
			w := wordList[qid]
			for i := range w {
				k := w[0:i] + "*" + w[i+1:]
				for _, wid := range wdict[k] {
					if wordList[wid] == endWord {
						return count
					}
					if !boolArr[wid] {
						boolArr[wid] = true
						nextQ = append(nextQ, wid)
					}
				}
			}
		}
		q = nextQ
	}
	return 0
}

func initWdict(wordList []string) map[string][]int {
	ret := map[string][]int{}
	for idx, w := range wordList {
		for i := range w {
			k := w[0:i] + "*" + w[i+1:]
			if _, ok := ret[k]; !ok {
				ret[k] = []int{}
			}
			ret[k] = append(ret[k], idx)
		}
	}
	return ret
}
