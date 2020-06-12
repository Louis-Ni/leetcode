package main

import "fmt"

func main() {
	que := [][]int{{1, 0}}
	ans := uniquePathsWithObstacles(que)
	fmt.Println(ans)
}
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	colLength := len(obstacleGrid)
	rowLength := len(obstacleGrid[0])

	res := make([][]int, colLength)
	for r := range res {
		res[r] = make([]int, rowLength)
	}

	for row := 0; row < rowLength; row++ {
		if obstacleGrid[0][row] != 1 {
			res[0][row] = 1
		}
	}

	for col := 0; col < colLength; col++ {
		if obstacleGrid[col][0] != 1 {
			res[col][0] = 1
		}
	}

	for col := 1; col < colLength; col++ {
		for row := 1; row < rowLength; row++ {
			if obstacleGrid[col][row] == 1 {
				res[col+1][row] = res[col+1][row-1]
				res[col][row+1] = res[col-1][row+1]
				continue
			}
			res[col][row] = res[col-1][row] + res[col][row-1]
		}
	}
	fmt.Println(res)
	return res[colLength-1][rowLength-1]
}
