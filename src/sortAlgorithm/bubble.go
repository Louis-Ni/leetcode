package bubble

func Bubble(nums []int) []int{
	for i := 0; i < len(nums) - 1; i++{
		for j := 0; j< len(nums) - 1 -i; j++{
			if nums[j] > nums[j+1] {
				temp := nums[j+1]
				nums[j+1] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums
}
