package main

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Peak() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	return x
}
func main() {
	nums := [][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}
	kthSmallest(nums, 8)
}

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix) - 1
	m := len(matrix[0]) - 1
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if h.Len() < k {
				heap.Push(h, matrix[i][j])
			} else {
				if matrix[i][j] < h.Peak().(int) {
					heap.Pop(h)
					heap.Push(h, matrix[i][j])
				}
			}
		}
	}

	ans := heap.Pop(h).(int)

	fmt.Println(ans)
	return ans
}
