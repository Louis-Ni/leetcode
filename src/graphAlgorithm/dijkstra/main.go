package main

import (
	"container/heap"
	"fmt"
)

func main() {
	ans := make(map[string]map[string]int)
	ans["A"] = map[string]int{"B": 5, "C": 1}
	ans["B"] = map[string]int{"A": 5, "C": 2, "D": 1}
	ans["C"] = map[string]int{"A": 1, "B": 2, "D": 4, "E": 8}
	ans["D"] = map[string]int{"B": 1, "C": 4, "D": 3, "F": 6}
	ans["E"] = map[string]int{"C": 8, "D": 3}
	ans["F"] = map[string]int{"D": 6}
	fmt.Println(ans["A"]["B"])
}

func dijkstra(graph map[string][]string, start string) {
	queue := []string{}
	priorityQueue := &NodeSlice{}
	heap.Init(priorityQueue) //初始化優先隊列
	node := Node{start, 0}
	heap.Push(priorityQueue, node) //把起始節點加入到優先隊列中
	isTouch := make(map[string]int)
	queue = append(queue, start)
	isTouch[start] = 1 //標記
	distance := 0
	for len(queue) > 0 {
		node := queue[0] //彈出第一個
		if len(queue) >= 1 {
			queue = queue[1:]
		}
		nodes := graph[node]
		for _, node := range nodes {
			if _, ok := isTouch[node]; !ok {
				queue = append(queue, node)
				isTouch[node] = 1
			}
		}
		fmt.Println(node)
	}
}

type Node struct {
	Ver    string
	Length int
}

//按照长度 从小到大排序
type NodeSlice []Node

//计算长度
func (s NodeSlice) Len() int {
	return len(s)
}

//交换元素
func (s NodeSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s NodeSlice) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return s[i].Length < s[j].Length
}

//推进一个节点
func (s *NodeSlice) Push(x interface{}) {
	*s = append(*s, x.(Node))
}

func (s *NodeSlice) Pop() interface{} {
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}
