package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

// Let's build this line for checking whether IntHeap implements the heap interface
var _ heap.Interface = (*IntHeap)(nil)

func (pq IntHeap) Len() int {
	return len(pq)
}

func (pq IntHeap) Less(i, j int) bool {
	return pq[i] < pq[j]
}

func (pq IntHeap) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *IntHeap) Push(x interface{}) {
	*pq = append(*pq, x.(int))
}

func (pq *IntHeap) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

// Adaption of the heap
func main() {
	h := &IntHeap{1, 4, 5, 6, 21, 8, 4, 2, 55, 13, 24, 12, 23}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 2)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}

}
