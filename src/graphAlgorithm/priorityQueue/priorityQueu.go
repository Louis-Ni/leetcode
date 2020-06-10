package main

type Node struct {
	Ver string
	Length int
}
//按照长度 从小到大排序
type NodeSlice []Node

//计算长度
func (s NodeSlice) Len() int{
	return len(s)
}

//交换元素
func (s NodeSlice) Swap(i, j int){
	s[i], s[j] = s[j], s[i]
}

func (s NodeSlice) Less(i, j int) bool { // 重写 Less() 方法， 从小到大排序
	return s[i].Length < s[j].Length
}

//推进一个节点
func (s *NodeSlice) Push(x interface{}){
	*s = append(*s, x.(Node))
}

func (s *NodeSlice) Pop() interface{}{
	old := *s
	n := len(old)
	x := old[n-1]
	*s = old[0 : n-1]
	return x
}