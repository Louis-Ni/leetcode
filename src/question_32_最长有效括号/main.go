package main

import (
	"container/list"
)

/**

1.需有一个变量start记录有效括号子串的起始下标，max表示最长有效括号子串长度，初始值均为0

2.遍历给字符串中的所有字符

    2.1若当前字符s[index]为左括号'('，将当前字符下标index入栈（下标稍后有其他用处），处理下一字符

    2.2若当前字符s[index]为右括号')'，判断当前栈是否为空

        2.2.1若栈为空，则start = index + 1，处理下一字符（当前字符右括号下标不入栈）

        2.2.2若栈不为空，则出栈（由于仅左括号入栈，则出栈元素对应的字符一定为左括号，可与当前字符右括号配对），判断栈是否为空

            2.2.2.1若栈为空，则max = max(max, index-start+1)

            2.2.2.2若栈不为空，则max = max(max, index-栈顶元素值)
**/

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

func longestValidParentheses(s string) int {
	if len(s) <= 1 {
		return 0
	}
	stack := New()
	m := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if s[i] == 40 {
			stack.Push(i)
			continue
		} else {
			if stack.Len() == 0 {
				start = i + 1
				continue
			} else {
				stack.Pop()
				if stack.Len() == 0 {
					m = max(m, i-start+1)
				} else {
					m = max(m, i-stack.Top().(int))
				}
			}
		}
	}

	return m
}
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {

}
