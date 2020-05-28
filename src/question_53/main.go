package main

import "fmt"

func main() {
	arr := []int{-2, -1}
	max := maxSubArray(arr)
	fmt.Println(max)
}
func maxSubArray(nums []int) int {
	if len(nums) == 1{
		return nums[0]
	}
	max := nums[0]
	last := nums[0]
	for i:=1; i< len(nums); i++{
		last = Max(nums[i], last+nums[i])
		max = Max(last, max)
	}
	return max
}
func Max(x, y int) int{
	if x > y{
		return x
	}
	return y
}
