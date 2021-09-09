package isValidBrackets

import "fmt"

// 功能: 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
// 思路: 使用栈来处理, 来一个右括号,需要一个左括号来进行匹配, 如果无法进行匹配则不合法

func IsValid(s string) bool {
	mapping := map[uint8]uint8{
		'(': ')',
		'[': ']',
		'{': '}',
	}

	stack := ""
	for _, ch := range s {

		// 左括号, 直接入栈
		if ch == '(' || ch == '[' || ch == '{' {
			stack = fmt.Sprintf("%c%s", ch, stack)
		} else { // 右括号
			if len(stack) <= 0 || uint8(ch) != mapping[stack[0]] {
				return false
			}
			stack = stack[1:]
		}
	}

	// 如果还有未匹配的左括号, 则不合法
	if len(stack) != 0 {
		return false
	}

	return true
}
