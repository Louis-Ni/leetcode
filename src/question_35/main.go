package main

import "fmt"

func main() {
	nums := []int {1,3}
	ans := searchInsert(nums, 2)
	fmt.Println(ans)
}
func searchInsert(nums []int, target int) int {
	if len(nums) == 1 {
		if target == nums[0] {
			return 0
		} else if target <= nums[0] {
			return 0
		} else {
			return 1
		}
	}
	pos := binarySearch(nums, target)
	if pos != -1 {
		return pos
	} else {
		slow := 0
		fast := 1
		for i := 0; i < len(nums)-1; i++ {
			if nums[slow] >= target {
				return slow
			}
			if target >= nums[slow] && target <= nums[fast] {
				return i+1
			}
			slow++
			fast ++
		}
	}
	return len(nums)
}
func binarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else if nums[mid] > target {
			right = mid - 1
		}
	}
	return -1
}
