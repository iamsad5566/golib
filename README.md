# Go common library
[![Go Report Card](https://goreportcard.com/badge/github.com/iamsad5566/golib)](https://goreportcard.com/report/github.com/iamsad5566/golib)   
    A go-library includes several customized implementation data structure, keep updating.

## Version:
    v1.2.0
## Install:
    go get github.com/iamsad5566/golib

## Package list:
1. priorityQ
    1. Max Heap
        > All the elements within the heap are sorted in decreasing order.
    2. Min Heap
        > All the elements within the heap are sorted in increasing order.
    - Methods:
        - Push(i interface{})
            > add new element into the heap.
        - Pop(): interface{}
            > return the min/max element from the min/max heap and remove.
        - Len(): int
            > return the size of the heap.

2. set
   > Set stores the input elements, it's an unordered structure, and does not store duplicate elements
    - Methods:
        - Add(T)
        - Remove(): bool
        - Len(): int
        - Next(): int, err

3. slice
    > Slice includes a method, Contains so far
    - Methods:
        - Contains(T): bool
        > return a boolean value showing whether the slice contains the input value.

4. stack, queue
    > stack: FILO, queue: FIFO  
      
    stack methods:  
    - Push(T) 
    - Pop(): T
    - IsEmpty(): bool
    - Len(): int

    queue methods:  
    - Add(T)
    - Poll(): T
    - IsEmpty(): bool
    - Len(): int
