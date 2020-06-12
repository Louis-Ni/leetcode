package main

import "fmt"

func main() {
	ans := climbStairs(4)
	fmt.Println(ans)
}
func climbStairs(n int) int {
	solutionOne := 1
	solutionTwo := 2
	if n == 1{
		return solutionOne
	}
	if n ==2 {
		return solutionTwo
	}
	answer:=0
	for i:=3; i<=n; i++{
		answer = solutionTwo + solutionOne
		solutionOne = solutionTwo
		solutionTwo = answer
	}
	return answer
}