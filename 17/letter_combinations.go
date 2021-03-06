package letterCombinations

import "fmt"

// 功能: 给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
// 思路: 每次取一个字符, 然后子串进行递归, 直到子串长度为0退出

func LetterCombinations(digits string) []string {

	var digitsMap = map[int][]int{
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	var results []string
	if len(digits) == 1 {
		list := digitsMap[int(digits[0])]
		for _, ch := range list {
			results = append(results, fmt.Sprintf("%c", ch))
		}
		return results
	}

	for lix, digit := range digits {
		list := digitsMap[int(digit)]
		for _, ch := range list {
			elements := LetterCombinations(digits[lix+1:])
			for _, element := range elements {
				result := fmt.Sprintf("%c%s", ch, element)
				if len(result) == len(digits) {
					results = append(results, result)
				}
			}
		}
	}
	return results
}
