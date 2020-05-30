package main

import "fmt"

func main() {
	nums := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ans:=maxArea(nums)
	fmt.Println(ans)
}
func maxArea(height []int) int {
	i,j:=0, len(height)-1
	water := 0
	for i<j{
		water = max (water, (j-i)* min(height[i], height[j]))//当前最大面积 和 正在计算的面积进行对比
		if height[i] < height[j]{
			i +=1
		}else{
			j-=1
		}
	}
	return water
}
func max(x,y int)int{
	if x> y{
		return x
	}
	return y
}
func min(x, y int)int{
	if x<y{
		return x
	}
	return y
}