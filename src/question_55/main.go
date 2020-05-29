package main

const LENGTH int = 0
func main() {
	nums := []int{3,1,0,2,4}
	canJump(nums)
}
func canJump(nums []int) bool {
	var dp = make([]int, len(nums))
	for i:=0; i<len(nums); i++ {
		dp[i]  = 0
	}
	dp[len(nums)-1] = 1
	ans:= recurrentJump(0, nums, &dp)
	return ans
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
