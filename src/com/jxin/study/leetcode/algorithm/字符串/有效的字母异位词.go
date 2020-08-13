package 字符串

import "sync"

/**
异步执行
用信号量
*/
func isAnagram1(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	semaphore := sync.WaitGroup{}
	semaphore.Add(2)
	sArr := [26]int{}
	tArr := [26]int{}
	go func() {
		warpMap(&sArr, s)
		semaphore.Done()
	}()
	go func() {
		warpMap(&tArr, t)
		semaphore.Done()
	}()
	semaphore.Wait()
	return sArr == tArr
}

/**
异步执行
用 渠道
*/
func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sArr := <-async(s)
	tArr := <-async(t)
	return sArr == tArr
}
func async(s string) chan [26]int {
	ret := make(chan [26]int, 1)
	go func() {
		sArr := [26]int{}
		warpMap(&sArr, s)
		ret <- sArr
	}()
	return ret
}
func warpMap(m *[26]int, s string) {
	for _, v := range s {
		idx := v - 'a'
		m[idx]++
	}
}
