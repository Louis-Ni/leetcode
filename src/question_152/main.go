package main

func main() {

}
func maxProduct(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	small := nums[0]
	large := nums[0]
	global := large
	for i := 1; i < len(nums); i++ {
		temp := large * nums[i]
		large = max(nums[i], max(large*nums[i], small*nums[i]))
		small = min(nums[i], min(temp, small*nums[i]))
		global = max(global, large)
	}
	return global
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
