package main

import "fmt"

func main() {
	res := integerReplacement(7)
	fmt.Println(res)
}
func integerReplacement(n int) int {
	return getNum(n,0)
}

func getNum(n int, num int) int{
	if n == 1{
		return num
	}
	if n % 2 == 0{
		//偶数
		num++
		return getNum(n/2, num)
	}else{
		//奇数
		//当n是奇数的时候，我们取n-1，和n+1计算次数的最小值即可。
		num++
		return min(getNum(n+1, num),getNum(n-1, num))
	}
}

func min(x , y int ) int {
	if x < y{
		return x
	}
	return y
}