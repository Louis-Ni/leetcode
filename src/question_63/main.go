package main

import "fmt"

func main() {
	que := [][]int{{0,1}}
	ans := uniquePathsWithObstacles(que)
	fmt.Println(ans)
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m := len(obstacleGrid)
	n := len(obstacleGrid[0])
	res := make([][]int, m)

	for r := range res {
		res[r] = make([]int, n)
	}

	path := 0
	if obstacleGrid[0][0] == 1 {
		path = 0
	} else {
		path = 1
	}

	for row := 1; row < n && obstacleGrid[0][row] != 1; row++ {
		if res[0][row] == 1 {
			path = 0
		}
		res[0][row] = path
	}

	path = 1
	for col := 0; col < m && obstacleGrid[col][0] != 1; col++ {
		if res[col][0] == 1 {
			path = 0
		}
		res[col][0] = path
	}

	for col := 1; col < m; col++ {
		for row := 1; row < n; row++ {
			fmt.Println("into")
			if obstacleGrid[col][row] == 1 {
				res[col][row] = 0
			} else {
				res[col][row] = res[col][row-1] + res[col-1][row]
			}
		}
	}
	fmt.Println(res)
	return res[m-1][n-1]
}
