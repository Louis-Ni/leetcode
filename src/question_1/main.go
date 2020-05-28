package main

func main() {
	nums :=[]int{2,7,11,15}
	target:=9
	twoSum(nums, target)
}
func twoSum(nums []int, target int) []int {
	count := make(map[int]int)
	n := []int{0}
	for i:=0; i< len(nums); i++ {
		var sub = target - nums[i]
		v1,ok := count[sub]
		if  ok {
			m := []int{v1,i}
			return m
		}else{
			num := nums[i]
			count[num] = i
		}
	}
	return n
}