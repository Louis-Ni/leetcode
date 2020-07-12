package main

import "fmt"

func main() {
	obj := Constructor()
	// obj.q.Push(1)
	q := Stack{0, []int{}}
	q.Push(1)
	fmt.Println()
	fmt.Println(obj)
}

type Stack struct {
	i    int //下標
	data []int
}

func (s *Stack) Push(k int) {
	s.data[s.i] = k
	s.i++
}

func (s *Stack) Pop() (ret int) {
	s.i--
	ret = s.data[s.i]
	return
}
func (s *Stack) Len() int {
	return s.i
}

type CQueue struct {
	p *Stack //保存數據
	q *Stack //保存臨時數據
}

func Constructor() CQueue {
	p := new(Stack)
	q := new(Stack)
	queue := CQueue{p, q}
	return queue
}

func (this *CQueue) AppendTail(value int) {
	if this.p.Len() != 0 {
		for this.p.Len() > 0 {
			this.q.Push(this.p.Pop())
		}
	}
	this.p.Push(value)
}

func (this *CQueue) DeleteHead() int {
	return 0
}

/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */
