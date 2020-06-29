package main

import (
	"container/heap"
	"fmt"
	"sort"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool{
	return h[i]<h[j]
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


func findKthLargest01(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sort.Ints(nums)
	h := &IntHeap{}
	heap.Init(h)
	length :=len(nums) - 1
	for k > 0{
		heap.Push(h, nums[length])
		length --
		k --
	}
	return  heap.Pop(h).(int)
}
func findKthLargest(nums []int, k int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	sort.Ints(nums)
	fmt.Println(nums)
	count := k
	length := len(nums) - 1
	ans := 0
	for count > 0 {
		ans = nums[length]
		length--
		count--
	}
	return ans
}
func main() {
	nums := []int{3,2,1,5,6,4}
	ans := findKthLargest01(nums, 2)
	fmt.Println(ans)
}