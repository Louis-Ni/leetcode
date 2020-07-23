package main

import "fmt"

func main() {
	nums := [][]int{{1,3,1},{1,5,1},{4,2,1}}
	res:=minPathSum(nums)
	fmt.Println(res)
}
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	row := len(grid)
	col := len(grid[0])
	var dp [][]int
	//初始化数组
	for i:=0; i<row; i++{
		arr := make([]int, len(grid[0]))
		dp = append(dp, arr)
	}
	dp[row-1][col-1] = grid[row-1][col-1]

	//最后一列
	for i:= row-2; i>=0; i--{
		dp[i][col-1] = dp[i+1][col-1] + grid[i][col-1]
	}

	//最后一行
	for i:=col-2; i>=0; i--{
		dp[row-1][i] = dp[row-1][i+1]+grid[row-1][i]
	}

	for i := row - 2; i >= 0; i--{
		for j := col - 2; j >= 0; j--{
			dp[i][j] = min(dp[i+1][j], dp[i][j+1]) + grid[i][j]
		}
	}

	fmt.Println(dp)
	return dp[0][0]
}

func min(x, y int) int{
	if x < y{
		return x
	}
	return y
}