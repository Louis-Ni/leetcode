package priorityQueue

type graph struct {
	vert string
	edge *[]edge
}

type edge struct {
	e string
	p int
}
type Graph []graph

func (g *Graph) Len() int{ //获取队列长度
	return len(*g)
}

func (g *Graph) Less(i,j int) bool{ //比较元素大小
	return (*g)[i].length < (*g)[j].length
}

func (g *Graph) Swap(i, j int){
	(*g)[i] , (*g)[j] = (*g)[j], (*g)[i]
}

func (g *Graph) Push(x interface{}){
	*g = append(*g, x.(graph))
}

func (g *Graph) Pop() interface{} {
	res := (*g)[0]
	*g = (*g)[1:]
	return  res
}