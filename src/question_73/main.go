package main

func main() {
	nums := [][]int{{0,1,2,0},{3,4,5,2},{1,3,1,5}}
	setZeroes(nums)
}
func setZeroes(matrix [][]int)  {
	colLength := len(matrix)
	rowLength := len(matrix[0])
	isCol := false
	isRow := false
	// 记录第一列为0的情况
	for c:= 0; c<colLength; c++{
		if matrix[c][0] == 0{
			isCol = true
			break
		}
	}

	//记录第一行有0的情况
	for r:=0; r<rowLength;r++{
		if matrix[0][r] == 0{
			isRow = true
			break
		}
	}
	// 从第1行第1列开始查。如果某个数字是0 就把他当前列和行的第0个设置为0
	for i:=1; i<colLength; i++{
		for j:=1; j<rowLength; j++{
			if matrix[i][j] == 0{
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 查找某个被设置成为0的的第一行或第一列。并且把该数值也设置成为0
	for i:=1; i<colLength; i++{
		for j:=1; j<rowLength; j++{
			if matrix[i][0] == 0 || matrix[0][j] == 0{
				matrix[i][j] = 0
			}
		}
	}
	//将所有第一列设置为0
	if isCol == true{
		for c:= 0; c<colLength; c++{
			matrix[c][0] = 0
		}
	}

	//将第一行设置为0
	if isRow == true{
		for r:=0; r<rowLength;r++{
			 matrix[0][r] = 0
		}
	}
}