package main

import "fmt"

func main() {
	nums := []int{1,3,5}
	res := minArray(nums)
	fmt.Println(res)
}
func minArray(numbers []int) int {
	if len(numbers) == 1{
		return numbers[0]
	}
	start:=0

	res := numbers[0]
	for i:=1; i<len(numbers); i++{
		if numbers[i] >= numbers[start]{
			start++
			continue
		}
		res = numbers[i]
		break
	}
	return res
}