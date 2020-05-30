package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-0,2,1,-3}
	ans:=threeSumClosest(nums, 1)
	fmt.Println(ans)
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	closest := nums[0] + nums[1] + nums[2]
	diff := cabs(closest - target)
	for i:= 0; i< len(nums) - 2; i++{
		if i!= 0 && nums[i] == nums[i-1]{
			continue//为了不重复做同一个数字
		}
		left := i+1
		right := len(nums)-1
		for left < right{
			sum :=  nums[i] + nums[left] + nums[right]
			fmt.Printf("%d+%d+%d=%d\n", nums[i] , nums[left] , nums[right], sum)
			newDiff := cabs(sum-target)
			if diff > newDiff{
				diff = newDiff
				closest = sum
			}
			if sum < target{
				left += 1
			}else{
				right -=1
			}
		}
	}
	return closest
}
func cabs(x int) int{
	if x < 0{
		return x * (-1)
	}
	return x
}
func max (x,y int) int{
	if x> y{
		return x
	}
	return y
}
func min(x,y int) int {
	if x < y {
		return x
	}
	return y
}