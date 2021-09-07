package longestCommonPrex

import "fmt"

// 功能: 求最长公共子串
// 思路: 取所有字符串最短的, 然后挨个取所有字符串的每个字符，遇到有不一样的退出, 得到最长公共子串

func LongestCommonPrefix(strs []string) string {

	// 获取最短长度
	minLength := len(strs[0])
	for _, str := range strs {
		length := len(str)
		if length < minLength {
			minLength = length
		}
	}

	result := ""
	for i := 0; i < minLength; i++ {
		ch := strs[0][i]
		for _, str := range strs {
			if str[i] != ch {
				return result
			}
		}
		result = fmt.Sprintf("%s%c", result, ch)
	}
	return result
}
