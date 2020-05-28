package main

import "fmt"
import "sortAlgorithm"

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	sortNums := bubble.Bubble(nums)
	result := threeSum(sortNums)
	fmt.Println(sortNums)
	fmt.Println(result)
}
func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		if i == 0 || nums[i] != nums[i-1] {
			start := i + 1
			end := len(nums) - 1
			for start < end {
				if nums[i]+nums[start]+nums[end] == 0 {
					ans = append(ans, []int{nums[i], nums[start], nums[end]})
					start += 1
					end -= 1
					for start < end && nums[start] == nums[start-1] {
						start += 1
					}
					for start < end && end+1 < len(nums) && nums[end] == nums[end+1] {
						end -= 1
					}
				} else if nums[i]+nums[start]+nums[end] < 0 {
					start += 1
				} else {
					//大于0的情况下
					end -= 1
				}

			}
		}
	}
	return ans
}
