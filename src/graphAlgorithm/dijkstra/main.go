package main

import (
	"container/heap"
	"fmt"
)


func main() {
	ans := make(map[string]map[string]int)
	ans["A"] =  map[string]int{"B":5,"C":1}
	ans["B"] = map[string]int{"A":5, "C":2, "D":1}
	ans["C"] = map[string]int{"A":1, "B":2, "D":4, "E":8}
	ans["D"] = map[string]int{"B":1, "C":4, "D":3, "F":6}
	ans["E"] = map[string]int{"C":8, "D":3}
	ans["F"] = map[string]int{"D":6}
	heap.Init(ans)
}

func dijkstra(graph map[string][]string, start string) {
	queue := []string{}
	isTouch := make(map[string]int)
	queue = append(queue, start)
	isTouch[start] = 1 //標記

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