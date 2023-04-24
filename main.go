package main

import (
	"container/heap"
	"fmt"
)

// Adaption of the heap
func main() {
	h := &MaxHeap{1, 4, 5, 6, 21, 8, 4, 2, 55, 13, 24, 12, 23}
	heap.Init(h)
	heap.Push(h, 3)
	heap.Push(h, 2)
	for h.Len() > 0 {
		fmt.Println(heap.Pop(h))
	}

	fmt.Println("=======================")

	hMin := &MinHeap{1, 4, 5, 6, 21, 8, 4, 2, 55, 13, 24, 12, 23}
	heap.Init(hMin)
	for hMin.Len() > 0 {
		fmt.Println(heap.Pop(hMin))
	}
}
