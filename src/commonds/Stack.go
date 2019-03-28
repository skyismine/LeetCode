package commonds

import "errors"

/**
栈结构和实现

详细文档参见 doc/commonds/Stack.md
*/
type Stack struct {
	// 栈存储空间
	element []interface{}
	// 栈顶 index
	top int
}

/**
创建 Stack
 */
func NewStack() *Stack {
	return &Stack{make([]interface{}, 0), 0}
}


// 入栈
func (stack *Stack) Push(e interface{}) error {
	stack.element = append(stack.element, e)
	stack.top = len(stack.element)-1
	return nil
}

// 出栈
func (stack *Stack) Pop() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}
	e := stack.element[stack.top]
	stack.element = stack.element[:stack.top]
	stack.top--

	return e, nil
}

// 栈是否为空
func (stack *Stack) IsEmpty() bool {
	return len(stack.element) == 0
}

//　返回栈顶元素但不删除它
func (stack *Stack) Top() (interface{}, error) {
	if stack.IsEmpty() {
		return nil, errors.New("stack is empty")
	}

	return stack.element[stack.top], nil
}