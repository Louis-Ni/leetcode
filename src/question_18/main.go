package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-3,-2,-1,0,0,1,2,3}
	ans:=fourSum(nums, 0)
	fmt.Println(ans)
}
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	ans := make([][]int, 0)
	for i:=0; i< len(nums)-3; i++{
		if i != 0 && nums[i] == nums[i-1]{
			continue //为了不重复做同一个数字
		}
		for j:= i+1; j< len(nums) - 2; j++{
			if j!= i+1 && nums[j] == nums[j-1]{
				continue//为了不重复做同一个数字
			}
			left := j+1
			right := len(nums)-1
			for left < right{
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum < target{
					left += 1
				}else if sum > target{
					right -= 1
				}else{
					ans = append(ans, []int{nums[i], nums[j], nums[left], nums[right]})
					left += 1
					right -=1
					for left < right && nums[left] == nums[left-1]{
						left += 1
					}
					for  left < right && nums[right] == nums[right+1]{
						right -= 1
					}
				}
			}
		}
	}
	return  ans
}