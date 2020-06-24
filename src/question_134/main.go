package main

func main() {

}

func canCompleteCircuit(gas []int, cost []int) int {
	totalStation := 0
	totalCost := 0
	for i := 0; i < len(gas); i++ {
		totalStation += gas[i]
		totalCost += cost[i]
	}
	if totalStation-totalCost < 0 { // 说明所有油站的油加起来 都不够路途的消耗
		return -1
	}

	currentGas := 0
	start := 0
	/**
	这一个for里面是存在一个假设的
	假设 从 i -> i+1 可以走, 那么意味着 当前的剩余的油一定是 >=0
	而我们已经排除掉了 <0 的条件了, 因此在这个环当中 一定存在某个起点 可以绕完整个圈
	假设 i0 -> i1 可以 剩余油耗>=0
		i1 -> i2 可以 剩余油耗>=0
		i2 -> i3 可以 剩余油耗>=0
		i3 -> i4 不可以 那就是说 i3->i4这个路程的剩余油量 < 0
	那就从i4这个点继续向前, 直到最后一个顶点, 因为我们的假设是 一定存在某个点的油耗是0, 而且能从该顶点绕一圈
	*/
	for j := 0; j < len(gas)-1; j++ {
		currentGas = currentGas - cost[j] + gas[j] // 去往下一个油站的油耗
		if currentGas < 0 {                        // 如果不够 则从 第i + 1个继续往后走
			currentGas = 0 //设置油耗为0
			start = j + 1  //从下一个开始
		}
	}

	return start
}
