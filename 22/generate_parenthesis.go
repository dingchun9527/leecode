package generateParenthesis

// 功能: 数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
// 思路: 左括号小于等于右括号, 递归尝试所有可能性

func generateParenthesis(n int) []string {

	return getParenthesis("", n, n)
}

func getParenthesis(content string, left, right int) []string {
	if left == 0 && right == 0 {
		return []string{content}
	}

	// 左边括号等于右边, 此时只能用左括号
	var results []string
	if left == right {
		return getParenthesis(content+"(", left-1, right)
	} else {
		if left > 0 {
			result := getParenthesis(content+"(", left-1, right)
			results = append(results, result...)
		}
		result := getParenthesis(content+")", left, right-1)
		results = append(results, result...)
	}
	return results
}
