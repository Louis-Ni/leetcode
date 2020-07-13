package main

import "fmt"

func main() {
	A := []int{1, 1, 0, 0, 1}
	B := []int{1, 1, 0, 0, 0}

	res := findLength(A, B)
	fmt.Println(res)
}
func findLength(A []int, B []int) int {
	m := len(A)
	n := len(B)
	res := 0
	//A定 B移
	for i := 0; i <= n; i++ {
		windowLen := min(m, n-i)
		maxLen := getMaxLen(A, B, 0, i, windowLen)
		res = max(res, maxLen)
	}
	// B定 A移
	for j := 0; j <= m; j++ {
		windowLen := min(n, m-j)
		maxLen := getMaxLen(A, B, j, 0, windowLen)
		res = max(res, maxLen)
	}
	return res
}

func getMaxLen(A []int, B []int, aStart, bStart int, windowLen int) int {
	res := 0
	subLen := 0

	for i := 0; i < windowLen; i++ {
		if A[aStart+i] == B[bStart+i] {
			subLen++
		} else {
			subLen = 0
		}
		res = max(res, subLen)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
