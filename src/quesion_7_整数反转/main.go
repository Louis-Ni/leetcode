package main

import "math"
func main()  {
	
}

func reverse(x int)int {
	// 考虑第一个符号是负号的情况
	// 考虑最后一位数是0的情况
	// 考虑反转后数值> int32的情况
	// 考虑反转后数值< int32的情况

	ans := transform(x)

	if ans > math.MaxInt32 || ans < math.MinInt32{
		return 0
	}

	return ans
}


func transform(x int) int {
		ans := 0

		for x != 0{
			ans = ans * 10 + (x%10)
			x = x/10
		}

		return ans
}