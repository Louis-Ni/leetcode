package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{2, 1}
	ans := findKthLargest(nums, 2)
	fmt.Println(ans)
}

func findKthLargest(nums []int, k int) int {
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
