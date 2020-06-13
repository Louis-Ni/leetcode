package main

import "fmt"

func main() {
	nums := [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}}
	setZeroes(nums)
}
func setZeroes(matrix [][]int)  {
	colLength := len(matrix)
	rowLength := len(matrix[0])
	res := make([][]int, colLength)
	for a:=range res{
		res[a] = append([]int{}, matrix[a]...)
	}

	for i:=0; i<colLength; i++{
		for j:=0; j<rowLength; j++{
			if matrix[i][j] == 0{
				fmt.Printf("i:%d,j:%d,num:%d\n",i,j,matrix[i][j])
				for r:=0; r<rowLength; r++{
					res[i][r] = 0
				}
				for c:=0; c< colLength; c++{
					res[c][j] = 0
				}
			}
		}
	}
	for m := range matrix{
		matrix[m] = res[m]
	}
}