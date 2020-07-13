package main

import "fmt"

func main() {
	res := divingBoard(10, 15, 4)
	fmt.Println(res)
}

func divingBoard(shorter int, longer int, k int) []int {
	res := []int{}
	if k == 0 {
		return res
	}
	if shorter == longer {
		res = append(res, k*shorter)
		return res
	}
	res = append(res, k+1)
	for i := 0; i <= k; i++ {
		fmt.Println(longer*i + shorter*(k-i))
		res = append(res, longer*i+shorter*(k-i))
	}
	return res
}
