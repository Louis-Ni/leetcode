package main

func main() {

}

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	buyday := prices[0] //先設置購買日 也就是 最小買入值為 元素第一個值
	saleday := 0        // 設置賣出值為 0 也就是盈利最大為0

	for i := 0; i < len(prices); i++ {
		if prices[i] < buyday { // 如果元素第i個小於 買入的值 就更新它
			buyday = prices[i] // 更新最小買入值
		}
		if prices[i]-buyday > saleday { // 如果當前的值 - 買入的值 > 最大的利潤值
			saleday = prices[i] - buyday // 更新最大利潤值
		}
	}
	return saleday
}
