package 堆

import (
	"errors"
	"fmt"
)

type IHead interface {
	Add(val int) error
	Pop() (int, error)
	Delete(idx int) (int, error)
	IsEmpty() bool
	IsFull() bool
	Dump()
}
type ArrTreeHeap struct {
	heapSize  int
	nodeCount int
	heap      []int
}

func NewArrTreeHeap(heap []int) IHead {
	return &ArrTreeHeap{heapSize: 0, nodeCount: 3, heap: heap}
}
func (s *ArrTreeHeap) Add(val int) error {
	if s.IsFull() {
		return errors.New("堆已满")
	}

	s.heap[s.heapSize] = val
	s.changeParent(s.heapSize)
	s.heapSize++
	return nil
}
func (s *ArrTreeHeap) Pop() (int, error) {
	return s.Delete(0)
}
func (s *ArrTreeHeap) Delete(idx int) (int, error) {
	if s.IsEmpty() {
		return -1, errors.New("堆已空")
	}
	if idx > s.heapSize {
		return -1, errors.New("要删的值越界")
	}
	if idx == s.heapSize {
		result := s.heap[idx]
		s.heapSize--
		return result, nil
	}

	result := s.heap[idx]
	s.heap[idx] = s.heap[s.heapSize-1]
	s.heapSize--
	s.changeChild(idx)
	return result, nil
}

func (s *ArrTreeHeap) IsEmpty() bool {
	return s.heapSize == 0
}
func (s *ArrTreeHeap) IsFull() bool {
	return s.heapSize == len(s.heap)
}

func (s *ArrTreeHeap) Dump() {
	for _, v := range s.heap[0:s.heapSize] {
		fmt.Printf("%d ", v)
	}
}

func (s *ArrTreeHeap) changeParent(idx int) {
	if idx == 0 {
		return
	}
	parentIdx := (idx - 1) / s.nodeCount
	parentVal := s.parent(parentIdx)
	node := s.heap[idx]
	if node < parentVal {
		return
	}
	s.heap[idx], s.heap[parentIdx] = parentVal, node
	s.changeChild(parentIdx)
}
func (s *ArrTreeHeap) changeChild(idx int) {
	leftIdx := idx*s.nodeCount + 1
	rightIdx := idx*s.nodeCount + 2
	leftVal := s.child(leftIdx)
	rightVal := s.child(rightIdx)
	node := s.heap[idx]
	if leftVal > rightVal {
		if node >= leftVal {
			return
		}
		s.heap[idx], s.heap[leftIdx] = leftVal, node
		s.changeChild(leftIdx)
		return
	}

	if rightVal == -1 || node >= rightVal {
		return
	}
	s.heap[idx], s.heap[rightIdx] = rightVal, node
	s.changeChild(rightIdx)
	return
}
func (s *ArrTreeHeap) parent(idx int) int {
	if idx < 0 {
		return -1
	}
	return s.heap[idx]
}
func (s *ArrTreeHeap) child(idx int) int {
	if idx > s.heapSize {
		return -1
	}
	return s.heap[idx]
}
