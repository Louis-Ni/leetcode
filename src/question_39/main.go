package main

import "fmt"

func main() {
	nums := []int{2, 3, 6, 7}
	target := 7
	ans := combinationSum(nums, target)
	fmt.Println(ans)
}

func combinationSum(candidates []int, target int) [][]int {
	res := make([][]int,0)
	tmp := []int{}
	recursiveHelper(&res, tmp, candidates, target, 0)
	return res
}

func recursiveHelper(res *[][]int, tmp []int, candidates []int, target int, index int) {
	if target <= 0 {
		if target == 0 {
			*res = append(*res, append([]int{}, tmp...)) //此处一定要新建一个数组重新赋值
		}
		return
	}
	for i := index; i < len(candidates); i++ {
		tmp = append(tmp, candidates[i])
		recursiveHelper(res, tmp, candidates, target-candidates[i], i)
		tmp = tmp[:len(tmp)-1]
	}
}
