package main

import "fmt"

func main() {
	ans := uniquePaths(3, 7)
	fmt.Println(ans)
}
func uniquePaths(m int, n int) int {
	res := make([][]int, m)
	for r := range res {
		res[r] = make([]int, n)
	}

	for row := 0; row < n; row++ {
		res[0][row] = 1
	}
	for col := 0; col < m; col++ {
		res[col][0] = 1
	}

	for row := 1; row < n; row++ {
		for col := 1; col < m; col++ {
			res[col][row] = res[col][row-1] + res[col-1][row]
		}
	}
	return res[m-1][n-1]
}
