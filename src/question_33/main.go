package main

import (
	"fmt"
)

func main() {
	nums := []int{4,5,6,7,0,1,2}
	ans:=search(nums, 1)
	fmt.Printf("ans = %d\n",ans)
}
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	left := 0
	right := len(nums) - 1
	for left + 1 < right{
		mid :=  (left + right) / 2
		fmt.Println(mid)
		if nums[mid] == target{
			return mid
		}else if nums[mid] < nums[right]{
			if nums[right] >= target && nums[mid]<=target{
				left = mid
			}else {
				right = mid
			}
		}else if nums[left] < nums[mid]{
			if nums[left] <= target && nums[mid] >= target{
				right = mid
			}else{
				left = mid
			}

		}
	}
	if nums[left] == target {
		return left
	}
	if nums[right] == target{
		return right
	}
	return -1
}