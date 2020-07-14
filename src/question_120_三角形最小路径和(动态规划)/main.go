package main

import "fmt"

func main() {
	nums := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
	// nums := [][]int{{-1}, {2, 3}, {1, -1, -3}}
	res := minimumTotal(nums)
	fmt.Println(res)
}

func minimumTotal(triangle [][]int) int {

	m := len(triangle)
	n := len(triangle[m-1])
	dp := make([]int, n)
	//底部最后一行的值
	for i := 0; i < len(dp); i++ {
		dp[i] = triangle[m-1][i]
	}

	//从倒数第二组开始算
	for i := m - 2; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}
