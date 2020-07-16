package main

import "fmt"

func main() {
	nums := [][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}}
	res := isBipartite(nums)
	fmt.Println(res)
}
func isBipartite(graph [][]int) bool {
	v := []int{}
	lenth := len(graph)

	for i := 0; i < lenth; i++ {
		v = append(v, -1)
	}
	for i := 0; i < lenth; i++ {
		if v[i] == -1 {
			queue := []int{}
			isTouch := make(map[int]int)

			queue = append(queue, i)
			isTouch[i] = 1
			v[i] = 0 //标记为0
			for len(queue) > 0 {
				node := queue[0]     //彈出第一個
				if len(queue) >= 1 { // 裁剪掉队列
					queue = queue[1:]
				}

				neighbor := 0
				if v[node] == 0 {
					neighbor = 1
				}

				nodes := graph[node]

				for _, n := range nodes {
					if v[n] == -1{ //如果未被标记
						queue = append(queue, n) //加入队列
						v[n] = neighbor //标记成为相反的
					}else if v[n] != neighbor{ //如果不一样 就说明这不是二分图
						return false
					}
				}
			}
		}
	}
	return true
}
