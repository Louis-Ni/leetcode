package main

import "fmt"

func main() {
	nums:= []int{2,7,9,11,15}
	target:=11
	res := twoSum(nums, target)
	fmt.Println(res)
}
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	var  res []int
	for i < j {
		if numbers[i] + numbers[j] == target{
			res = append(res, i +1)
			res = append(res, j +1)
			break
		}else if numbers[i] + numbers[j] < target{
			i++
			continue
		}else{
			j--
			continue
		}
	}

	return res
}