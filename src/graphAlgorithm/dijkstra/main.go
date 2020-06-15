package main

import (
	"container/heap"
	"fmt"
	"graphAlgorithm/priorityQueue"
)


func main() {
	ans := make(map[string]map[string]int)
	ans["A"] =  map[string]int{"B":5,"C":1}
	ans["B"] = map[string]int{"A":5, "C":2, "D":1}
	ans["C"] = map[string]int{"A":1, "B":2, "D":4, "E":8}
	ans["D"] = map[string]int{"B":1, "C":4, "D":3, "F":6}
	ans["E"] = map[string]int{"C":8, "D":3}
	ans["F"] = map[string]int{"D":6}
	dijkstra(ans,"A")
}
func initGraph(graph map[string]map[string]int, start string,dist *map[string]int){
	for w ,_ := range graph{
		if w != start{
			(*dist)[w] = 100000000000
		}
	}
}
func dijkstra(graph map[string]map[string]int, start string) {
	parent := make(map[string]string) //記錄某個節點的上一個節點
	isTouch := make(map[string]int) //看到過某個點的集合
	distance := make(map[string]int) //記錄頂點之間的距離
	initGraph(graph, start, &distance)
	fmt.Println(distance)
	pqueue := &priorityQueue.NodeSlice{}//優先隊列
	node := priorityQueue.Node{start, 0} //隊列中其中一個節點

	heap.Init(pqueue)//初始化優先隊列
	heap.Push(pqueue, node)//將首節點添加進去

	parent[start] = "None"

	for pqueue.Len() > 0 {
		vertNode := heap.Pop(pqueue).(priorityQueue.Node) //彈出第一個
		vert:=vertNode.Ver //頂點名稱
		length:=vertNode.Length //到頂點的距離
		isTouch[vert] = 1 //拿出來以後才標記成為已經處理過了
		nodes := graph[vert] //此處拿出來的是一個map 像這樣map[B:5 C:1]
		for w, _ := range nodes {
			if _, ok := isTouch[w]; !ok {
				if length + graph[vert][w] < distance[w]{
					tmpNode := priorityQueue.Node{w,length + graph[vert][w]}
					heap.Push(pqueue, tmpNode)
					parent[w] = vert
					distance[w] = length + graph[vert][w]

				}
			}
		}
	}
	fmt.Println(parent)
	fmt.Println(distance)
}