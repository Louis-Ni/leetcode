package Stack

import (
	"container/list"
)

type Stack struct {
	list *list.List
}

func New() *Stack {
	list := list.New()
	return &Stack{
		list: list,
	}
}

func (p *Stack) Push(v interface{}) {
	p.list.PushBack(v)
}

func (p *Stack) Pop() interface{} {
	e := p.list.Back()
	if e != nil {
		p.list.Remove(e)
		return e.Value
	}
	return nil
}

func (p *Stack) Top() interface{} {
	e := p.list.Back()
	if e != nil {
		return e.Value
	}
	return nil
}

func (p *Stack) Len() int {
	return p.list.Len()
}

func (p *Stack) Empty() bool {
	return p.list.Len() == 0
}
