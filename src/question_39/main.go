package main

import "fmt"

func main() {
	nums := []int{2, 3, 6, 7}
	target := 7
	ans := combinationSum(nums, target)
	fmt.Println(ans)
}

func combinationSum(candidates []int, target int) [][]int {
	res := [][]int{}
	tmp := []int{}
	recursiveHelper(&res, &tmp, candidates, target, 0)
	return res
}

func recursiveHelper(res *[][]int, tmp *[]int, candidates []int, target int, index int) {
	if target <= 0 {
		if target == 0 {
			fmt.Println("1")
			temp := (*tmp)
			fmt.Println(temp)
			(*res) = append((*res), temp)
			fmt.Println((*res))
		}
		return
	}
	for i := index; i < len(candidates); i++ {
		(*tmp) = append((*tmp), candidates[i])
		recursiveHelper(res, tmp, candidates, target-candidates[i], i)
		(*tmp) = (*tmp)[:len(*tmp)-1]
	}
}
