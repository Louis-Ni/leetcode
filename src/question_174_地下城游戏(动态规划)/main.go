package main

import "fmt"

func main() {
	num := [][]int{{1},{-2},{1}}
	res := calculateMinimumHP(num)
	fmt.Println(res)
}

func calculateMinimumHP(dungeon [][]int) int {
	if len(dungeon) == 0 || len(dungeon[0]) == 0 {
		return 0
	}

	row := len(dungeon)
	col := len(dungeon[0])
	var dp [][]int
	for i:=0; i<row; i++{
		arr := make([]int, len(dungeon[0]))
		dp = append(dp, arr)
	}

	m := -min(0,dungeon[row-1][col-1]) + 1
	dp[row-1][col-1] = m

	//先填最后一列
	for i := row - 2; i >= 0; i--{
		t := max(dp[i+1][col-1]-dungeon[i][col-1],0)
		if t == 0{
			dp[i][col-1] = 1
		}else{
			dp[i][col-1] = t
		}
	}
	//先填最后一行
	for i := col - 2; i >= 0; i--{
		t := max(dp[row-1][i+1]-dungeon[row-1][i],0)
		if t == 0{
			dp[row-1][i] = 1
		}else{
			dp[row-1][i] = t
		}
	}

	for i := row - 2; i >= 0; i--{
		for j := col - 2; j >= 0; j--{
			min := min(dp[i+1][j], dp[i][j+1])
			dp[i][j] = max((min-dungeon[i][j]),1)
		}
	}

	fmt.Println(dp)
	return dp[0][0]
}


func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
