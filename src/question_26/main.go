package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,2,2,2,3,4, 4, 5}
	ans:=removeDuplicates(nums)
	fmt.Println(ans)
}
func removeDuplicates(nums []int) int {
	if len(nums) <=1{
		return len(nums)
	}
	slow := 0
	for fast := 1; fast < len(nums); fast++{
		if nums[slow] != nums[fast]{
			if slow + 1 != fast{
				nums[slow+1] = nums[fast]
			}
			slow++
		}
	}
	nums = nums[:slow+1]
	return slow+1
}
func removeDuplicates2(nums []int) int {
	if len(nums) <= 1{
		return len(nums)
	}
	sort.Ints(nums)
	head := 0
	tail := 0
	for tail < len(nums) {
		fmt.Println(head, tail)
		if nums[head] == nums[tail] {
			tail ++
		} else {
			head++
			if head < tail {
				nums[head] = nums[tail]
				tail ++
			}
		}
	}
	fmt.Println(nums)
	nums = nums[:head+1]

	return head+1
}