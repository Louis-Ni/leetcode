package main

import "fmt"

func main() {
	ans := generateMatrix(3)
	fmt.Println(ans)
}
func generateMatrix(n int) [][]int {
	res := make([][]int, n)
	if n == 1{
		tmp := []int{1}
		res[0] = tmp
		return res
	}
	if n > 1{
		for a := range res{
			res[a] = make([]int, n)
		}
	}
	top := 0
	right := n - 1
	bottom := n - 1
	left := 0
	direction := "right"

	count := 1
	for left <= right && top <= bottom {
		switch direction {
		case "right":
			for r := left; r <= right; r++ {
				res[top][r] = count
				count++
			}
			top++ //做完后 整个跳到下一行
			direction = "bottom"
			break
		case "bottom":
			for b := top; b <= bottom; b++ {
				res[b][right] = count
				count++
			}
			right-- //做完后 整个右边往左靠一行
			direction = "left"
			break
		case "left":
			for l := right; l >= left; l-- {
				res[bottom][l] = count
				count++
			}
			bottom -- //做完后 整个下面往上一行
			direction = "top"
			break
		case "top":
			for t:= bottom; t>=top; t--{
				res[t][left] = count
				count++
			}
			left ++ //做完后 整个左边往右一行
			direction = "right"
			break
		}
	}
	return res
}
