package maxhp

import (
	"container/heap"
)

func InitializeMaxHeap() *maxHeap {
	hp := &maxHeap{}
	return hp
}

type maxHeap []int

// Let's build this line for checking whether MaxHeap implements the heap interface
var _ heap.Interface = (*maxHeap)(nil)

func (pq maxHeap) Len() int {
	return len(pq)
}

func (pq maxHeap) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq maxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *maxHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *maxHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}