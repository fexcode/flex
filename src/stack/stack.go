package stack

import (
	"errors"
)

// Stack 是一个可以存储任意类型的栈
type Stack struct {
	items []interface{}
}

// NewStack 创建一个新的栈
func NewStack() *Stack {
	return &Stack{
		items: make([]interface{}, 0),
	}
}

// Push 将一个元素压入栈顶
func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

// Pop 从栈顶弹出一个元素，并返回它
func (s *Stack) Pop() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

// Peek 返回栈顶元素，但不从栈中移除它
func (s *Stack) Peek() (interface{}, error) {
	if len(s.items) == 0 {
		return nil, errors.New("stack is empty")
	}
	return s.items[len(s.items)-1], nil
}

// Size 返回栈的大小
func (s *Stack) Size() int {
	return len(s.items)
}

// Data 返回栈中的所有元素
func (s *Stack) Data() []interface{} {
	return s.items
}

