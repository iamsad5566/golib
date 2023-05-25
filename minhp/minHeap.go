package minhp

import "container/heap"

func InitializeMinHeap() *MinHeap {
	hp := &MinHeap{}
	return hp
}

type MinHeap []int

// Let's build this line for checking whether MinHeap implements the heap interface
var _ heap.Interface = (*MinHeap)(nil)

func (pq MinHeap) Len() int {
	return len(pq)
}

func (pq MinHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq MinHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *MinHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *MinHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
