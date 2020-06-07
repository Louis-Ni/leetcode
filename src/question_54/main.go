package main

func main() {
	nums := [][]int{{1,2,3},{5, 6, 7, 8},{7,8,9}}
	spiralOrder(nums)
}
func spiralOrder(matrix [][]int) []int {
	res := []int{}
	if len(matrix) == 0{
		return res
	}
	top := 0
	right := len(matrix[0])-1
	bottom := len(matrix) - 1
	left := 0
	direction := "right"

	for left<=right && top <= bottom{
		switch direction {
		case "right":
			for r:=left; r<=right; r++{
				res = append(res, matrix[top][r])
			}
			top ++ //做完后 整个跳到下一行
			direction = "bottom"
			break
		case "bottom":
			for b:= top; b<=bottom; b++{
				res = append(res, matrix[b][right])
			}
			right -- //做完后 整个右边往左靠一行
			direction = "left"
			break
		case "left":
			for l:= right; l>=left; l--{
				res = append(res, matrix[bottom][l])
			}
			bottom -- //做完后 整个下面往上一行
			direction = "top"
			break
		case "top":
			for t:= bottom; t>=top; t--{
				res = append(res, matrix[t][left])
			}
			left ++ //做完后 整个左边往右一行
			direction = "right"
			break
		}
	}
	return res
}