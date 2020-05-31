package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1,3,4,8,1}
	nextPermutation(nums)
}
func nextPermutation(nums []int)  {
	n := len(nums)
	i:= n -2
	for i >= 0 && nums[i]>=nums[i+1] { //寻找一个最小可以交换位置的地方
		i--
	}
	if i>=0{
		j := n-1
		for j>i&&nums[j]<=nums[i]{ //寻找一个最小的值交换
			j--
		}
		swap(&nums, i, j)
	}
	sorted(&nums, i+1, n)
	fmt.Println(nums)
}
func sorted(nums *[]int, start, end int){
	temp := (*nums)[start:end]
	sort.Ints(temp)
	(*nums) = append((*nums)[:start], temp...)
}
func swap(nums *[]int, a, b int){
	temp := (*nums)[a]
	(*nums)[a] = (*nums)[b]
	(*nums)[b] = temp
}