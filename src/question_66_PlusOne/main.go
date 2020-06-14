package main

import "fmt"

func main() {
	nums := []int {9,9,9}
	plusOne(nums)
}
func plusOne(digits []int) []int {
	res := []int{0}
	carry :=0
	// 999
	lastOne := digits[len(digits)-1]
	if lastOne + 1 >= 10{ // 最后一位数 + 1 >= 10的时候 说明要进位
		carry = 1
		digits[len(digits)-1] = 0
		for i:=len(digits)-2; i>=0; i--{ //从倒数第二位开始数
			tmp := digits[i] + carry
			if tmp >= 10{ //如果倒数第i位开始 + carry 还>=10 说明倒数的第i位也需要进位
				carry = 1
				digits[i] = 0 //设置倒数第i为为0
			}else{
				if carry == 1{ //直到某一位不需要进位的时候， 该位 + 1 然后 再把carry 设置成为0
					digits[i] = digits[i] + 1
					carry = 0
				}
			}
		}
	}else{
		digits[len(digits)-1] = digits[len(digits)-1] + 1
	}
	if carry == 1 {
		res[0] = 1
		res = append(res, digits...)
	}else {
		res = digits
	}
	fmt.Println(res)
	return res
}
