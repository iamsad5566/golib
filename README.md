# Go common library
    A go-library includes several customized implementation data structure, keep updating.

## Version:
    v1.2.0
## Install:
    go get github.com/iamsad5566/golib

## Package list:
1. priorityQ
    1. Max Heap
    2. Min Heap
    - Methods:
        - Push(i interface{})
        - Pop(): interface{}
        - Len(): int

2. set
    - Methods:
        - Add(T)
        - Remove(): bool
        - Len(): int
        - Next(): int, err

3. slice
    - Methods:
        - Contains(T): bool

4. stack
