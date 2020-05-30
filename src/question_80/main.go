package main

import "fmt"

func main() {
	nums := []int{1,2,2,2,3,4,4,5}
	ans:=removeDuplicates(nums)
	fmt.Println(ans)
}

func removeDuplicates(nums []int) int {
	if len(nums) <=1{
		return len(nums)
	}
	slow := 0
	for fast := 2; fast < len(nums); fast++{
		if nums[slow] != nums[fast]{
			if slow + 2 != fast{
				nums[slow+2] = nums[fast]
			}
			slow++
		}
	}
	nums = nums[:slow+2]
	return slow+2
}