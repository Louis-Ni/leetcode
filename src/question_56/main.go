package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := [][]int{{1,4},{0,4},{8,10},{15,18}}
	merge(nums)
}
func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[j][0] > intervals[i][0]{
			return true
		}
		return false
	})
	fmt.Println(intervals)
	res := [][]int{}
	if len(intervals) < 1 {
		return res
	}
	current := intervals[0]
	res = append(res, current)
	for _, node := range intervals{
		if node[0] > current[1]{
			res = append(res, node)
			current = node
		}else{
			current[1] = max(current[1], node[1])
		}
	}
	return res
}
func max(x,y int) int{
	if x > y{
		return x
	}
	return y
}