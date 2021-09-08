package generateParenthesis

import "fmt"

// 功能: 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 思路: 使用递归的方式实现

func generateParenthesis(n int) []string {

	// 当n等于1, 达到递归边界, 直接返回()
	if n == 1 {
		return []string{"()"}
	}

	var results []string
	list := generateParenthesis(n-1)
	for _, item := range list{
		results = append(results, fmt.Sprintf("%s%s", item, "()"))
		results = append(results, fmt.Sprintf("%c%s%c", '(', item, ')'))
	}

	return results
}
