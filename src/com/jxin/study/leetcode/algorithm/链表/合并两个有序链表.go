package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

}
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	val1 := l1.Val
	val2 := l2.Val
	if val1 > val2 {
		next := l2.Next
		l2.Next, l1, l2 = l1, l2, next
	}
	resul := l1
	for l1 != nil && l2 != nil {
		next := l1.Next
		val := l2.Val
		if next == nil {
			l1.Next = l2
			return resul
		}
		if next.Val <= val {
			l1 = next
			continue
		}
		next2 := l2.Next
		l1.Next, l2.Next, l2 = l2, next, next2
	}
	return resul
}
