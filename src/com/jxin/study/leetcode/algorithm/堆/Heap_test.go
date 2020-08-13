package 堆

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	doTest(NewArrTreeHeap(make([]int, 20, 25)))
}

func doTest(head IHead) {
	head.Add(10)
	head.Add(4)
	head.Add(9)
	head.Add(1)
	head.Add(7)
	head.Add(5)
	head.Add(3)

	head.Dump()
	fmt.Printf("\n")
	head.Delete(5)
	head.Dump()
	fmt.Printf("\n")
	head.Pop()
	head.Dump()
	fmt.Printf("\n")
	head.Delete(2)
	head.Dump()
}
