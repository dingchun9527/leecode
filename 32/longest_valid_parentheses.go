package longestValidParentheses

import "fmt"

// 功能: 给你一个只包含 '(' 和 ')' 的字符串，找出最长有效（格式正确且连续）括号子串的长度。
// 思路: 使用栈来保存'(', 如果不足以弹出则栈清空

func longestValidParentheses(s string) int {
	stack, curr, max := "", 0, 0
	for _, ch := range s {
		if ch == '(' {	// 左括号, 入栈
			stack = fmt.Sprintf("%c%s", ch, stack)
		} else {	// 右括号, 弹栈
			if len(stack) <= 0 {	// 没有足够的左括号可弹出, 说明到这里括号已经不合法了
				curr = 0
			} else {
				stack = stack[1:]
				curr += 2
			}
			if curr > max && stack == "" {
				max = curr
			}
		}
	}

	return max
}
