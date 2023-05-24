package ykheap

import (
	"github.com/iamsad5566/ykheap/maxhp"
	"github.com/iamsad5566/ykheap/minhp"
)


func InitializeMaxHeap() *maxhp.MaxHeap {
	hp := &maxhp.MaxHeap{}
	return hp
}

func InitializeMinHeap() *minhp.MinHeap {
	hp := &minhp.MinHeap{}
	return hp
}