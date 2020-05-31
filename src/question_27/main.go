package main

import "fmt"

func main() {
	nums := []int{0,1,2,2,3,0,4,2}
	ans:=removeElement(nums, 2)
	fmt.Println(ans)
}
func removeElement(nums []int, val int) int {
	slow:=0
	for fast:=0; fast<len(nums); fast++{
		if nums[fast] != val{
			nums[slow] = nums[fast]
			slow++
		}
	}
	nums = nums[:slow]
	fmt.Println(nums)
	return slow
}