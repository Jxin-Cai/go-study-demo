package 高级搜索

import "testing"

func TestTrie(t *testing.T) {
	t.Log(Trie{})
}

type Trie struct {
	end  bool
	next map[int32]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{false, make(map[int32]*Trie)}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	if len(word) <= 0 {
		return
	}
	root := this
	for _, v := range word {
		if val, ok := root.next[v]; ok {
			root = val
		} else {
			next := &Trie{false, make(map[int32]*Trie)}
			root.next[v] = next
			root = next
		}
	}
	root.end = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	if len(word) <= 0 {
		return true
	}
	root := this
	for _, v := range word {
		if val, ok := root.next[v]; ok {
			root = val
		} else {
			return false
		}
	}
	return root.end
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	if len(prefix) <= 0 {
		return true
	}
	root := this
	for _, v := range prefix {
		if val, ok := root.next[v]; ok {
			root = val
		} else {
			return false
		}
	}
	return true
}
