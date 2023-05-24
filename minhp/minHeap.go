package minhp

import "container/heap"

func InitializeMaxHeap() *minHeap {
	hp := &minHeap{}
	return hp
}

type minHeap []int

// Let's build this line for checking whether MinHeap implements the heap interface
var _ heap.Interface = (*minHeap)(nil)

func (pq minHeap) Len() int {
	return len(pq)
}

func (pq minHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq minHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *minHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *minHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
