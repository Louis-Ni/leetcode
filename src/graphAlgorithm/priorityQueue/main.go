package main

import (
	"container/heap"
	"fmt"
)

func main() {
	ns := &NodeSlice{{"B",6},{"C",5},{"D",3},{"A",2}}
	heap.Init(ns)
	node := Node{"E",1}
	heap.Push(ns, node)

	for ns.Len() >0 {
		fmt.Println(heap.Pop(ns))
	}
}
