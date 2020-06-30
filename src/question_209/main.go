package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{5,1,3,5,10,7,4,9,2,8}
	ans := minSubArrayLen(15, nums)
	fmt.Println(ans)
}
func minSubArrayLen(s int, nums []int) int {
	if len(nums) == 0{
		return 0
	}
	sum, right, left, length := 0, 0, 0,math.MaxInt32
	for right<len(nums){
		sum += nums[right]
		right ++
		for sum >= s{
			length = min(length,right - left)
			sum -= nums[left]
			left++
		}
	}
	if length == math.MaxInt32{
		return 0
	}
	return length
}
func min(x,y int) int{
	if x < y{
		return x
	}
	return  y
}