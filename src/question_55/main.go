package main

import "fmt"

const LENGTH int = 0
func main() {
	nums := []int{3,2,1,0,4}
	canJump(nums)

}
func canJump(nums []int) bool {
	var dp = make([]int, len(nums))
	for i:=0; i<len(nums); i++ {
		dp[i]  = 0
	}
	dp[len(nums)-1] = 1
	//ans:= recurrentJump(0, nums, &dp)
	//ans := nonRecurrentJump(nums, &dp)
	ans:=optimizeNonRecurrentJump(nums)
	fmt.Println(ans)
	return ans
}
func optimizeNonRecurrentJump(nums[] int) bool{
	aim := len(nums)-1
	for i:= len(nums)-2;i>=0;i-- {
		maxJump := min(i+nums[i], len(nums)-1)
		for j := i + 1; j<=maxJump; j++{
			if j == aim{
				aim = i
			}
		}
	}
	if aim == 0{
		return true
	}
	return false
}
func nonRecurrentJump(nums[] int, dp *[]int) bool{
	for i:= len(nums)-2;i>=0;i-- {
		maxJump := min(i+nums[i], len(nums)-1)
		for j := i + 1; j<=maxJump; j++{
			if (*dp)[j] == 1{
				(*dp)[i] = 1
				break
			}
		}
	}
	if (*dp)[0] == 1{
		return true
	}
	return false
}
func recurrentJump(position int,nums []int, dp *[]int) bool{
	if (*dp)[position] == 1{
		return true
	}else if (*dp)[position] == -1{
		return false
	}
	maxJump := min(position + nums[position], len(nums)-1)
	for i:= position + 1; i<=maxJump; i++{
		result:=recurrentJump(i, nums, dp)
		if result {
			(*dp)[position] = 1
			return true
		}
	}
	(*dp)[position] = -1
	return false
}
func min(x int, y int) int{
	if (x < y){
		return x
	}
	return y
}
