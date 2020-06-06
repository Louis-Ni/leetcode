package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{10,1,2,7,6,1,5}
	target := 8
	ans := combinationSum2(nums, target)
	fmt.Println(ans)
}

func combinationSum2(candidates []int, target int) [][]int {
	res := make([][]int,0)
	tmp := []int{}
	count := 0
	sort.Ints(candidates)
	recursiveHelper(&res, tmp, candidates, target, 0, &count)
	fmt.Println(count)
	return res
}

func recursiveHelper(res *[][]int, tmp []int, candidates []int, target int, index int, count *int) {
	if target <= 0 {
		if target == 0 {
			*res = append(*res, append([]int{}, tmp...)) //此处一定要新建一个数组重新赋值
		}
		return
	}
	for i := index; i < len(candidates); i++ {
		if candidates[i] > target{
			return
		}
		if i>index&&candidates[i] == candidates[i-1]{
			continue
		}
		*count++
		tmp = append(tmp, candidates[i])
		recursiveHelper(res, tmp, candidates, target-candidates[i], i+1, count)
		tmp = tmp[:len(tmp)-1]
	}

}

