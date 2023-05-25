# Go common library
    A go-library includes several customized implementation data structure, keep updating.

## Version:
    v1.2.0
## Install:
    go get github.com/iamsad5566/golib

## Tool list:
1. priorityQ.Heap
    1. Max Heap
    2. Min Heap
    - Methods:
        - Push(i interface{})
        - Pop(): interface{}
        - Len(): int

2. set.Set
    - Methods:
        - Add(T)
        - Remove(): bool
        - Len(): int
        - Next(): int, err

3. slice.Contains
    - Methods:
        - Contains(T): bool
