package main

import "fmt"

func main() {
	ans := make(map[string][]string)
	ans["A"] = []string{"B", "C"}
	ans["B"] = []string{"A", "C", "D"}
	ans["C"] = []string{"A", "B", "D", "E"}
	ans["D"] = []string{"B", "C", "D", "F"}
	ans["E"] = []string{"C", "D"}
	ans["F"] = []string{"D"}
	dfs(ans, "E")
}
func dfs(graph map[string][]string, start string) {
	queue := []string{}
	isTouch := make(map[string]int)
	queue = append(queue, start)
	isTouch[start] = 1 //標記

	for len(queue) > 0 {
		n := len(queue)
		node := queue[n-1] //彈出第一個
		if len(queue) >= 1 {
			queue = queue[:n-1]
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
