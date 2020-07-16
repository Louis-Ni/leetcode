package main

import "fmt"

func main() {
	dic := []string{"looked","just","like","her","brother"}
	s:="jesslookedjustliketimherbrother"
	res := respace(dic, s)
	fmt.Println(res)
}
func respace(dictionary []string, sentence string) int {
	dp := make([]int, len(sentence)+1)
	right := 1
	for right <= len(sentence) {
		dp[right] = dp[right-1]
		for _, word := range dictionary { // 遍历字典内的值
			if right < len(word) {
				continue
			}
			//截取 sentence 的第 right - len(word) : right 的值
			//如果 word= like 那就是截取之前的4个字符 判断是否与 word相等
			tmp := sentence[right-len(word) : right]
			if tmp != word {
				continue
			}
			//相同的时候， 取最大值
			//从dp[right] 跟 dp[right-len(word)] + len(word) 对比 取最大
			//假设现在right= 11 字符串目前是 jesslookedj
			//word现在等于 looked
			//那就看dp[11]的值 和 dp[11-6]+6 的值谁大
			//就是说从dp[5]-dp[11]之间有可能匹配到某个比较小的字典 例如 ked 的值
			//那就要舍弃这个小的值，取最大匹配数的值
			dp[right] = max(dp[right], dp[right-len(word)] +len(word))

		}
		right ++
	}
	return len(sentence) - dp[len(sentence)]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
