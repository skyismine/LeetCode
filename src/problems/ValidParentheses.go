package problems

import (
	"commonds"
	"utils"
)

/**
题目描述：

给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
 */
func IsValid(s string) bool {
	switch METHOD {
	case METHOD_MINE:
		return isValidMine(s)
	case METHOD_LCA:
		return isValidLCA(s)
	}

	return false
}

func getParentheses(right byte) byte {
	if right == ')' {
		return '('
	} else if (right == ']') {
		return '['
	} else if (right == '}') {
		return '{'
	}

	return 0
}

/**
 使用栈来处理，将字符一个个进行处理，遇到　( [ { 则入栈，碰到　) ] } 则出栈比较
复杂度分析：
时间复杂度：O(n)，n 为字符串的长度
空间复杂度：O(n)，n 为字符串的长度，需要额外的 n 字符空间来构造栈结构
 */
func isValidMine(s string) bool {
	if len(s) == 0 {
		return true
	}

	sbytes := utils.StringBytes(s)
	stack := commonds.NewStack()
	for _,value := range sbytes {
		if value == '(' || value == '[' || value == '{' {
			err := stack.Push(value)
			if err != nil {
				return false
			}
		}
		if value == ')' || value == ']' || value == '}' {
			e, err := stack.Pop()
			if err != nil || e.(byte) != getParentheses(value) {
				return false
			}
		}
	}

	return stack.IsEmpty()
}

/**
 也是使用了栈，但用　map 保存括号对以简化代码
时间复杂度：O(n)，n 为字符串的长度
空间复杂度：O(n)，n 为字符串的长度，需要额外的 n 字符空间来构造栈结构
 */
func isValidLCA(s string) bool {
	pamap := map[byte]byte {
		')':'(',
		']':'[',
		'}':'{',
	}

	sbytes := utils.StringBytes(s)
	stack := commonds.NewStack()

	for _,value := range sbytes {
		if c,ok := pamap[value]; ok {
			e, err := stack.Pop()
			if err != nil || e.(byte) != c {
				return false
			}
		} else {
			stack.Push(value)
		}
	}
	return stack.IsEmpty()
}