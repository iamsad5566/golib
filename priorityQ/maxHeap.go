package priorityQ

import (
	"container/heap"
)

func InitializeMaxHeap() *MaxHeap {
	hp := &MaxHeap{}
	heap.Init(hp)
	return hp
}

type MaxHeap []int

// Let's build this line for checking whether MaxHeap implements the heap interface
var _ heap.Interface = (*MaxHeap)(nil)

func (pq MaxHeap) Len() int {
	return len(pq)
}

func (pq MaxHeap) Less(i, j int) bool {
	return pq[i] > pq[j]
}

func (pq MaxHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MaxHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *MaxHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
