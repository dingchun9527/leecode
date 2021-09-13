package findSubstring

import "strings"

// 思路: 先遍历words所能排列出地不同组合, 再挨个匹配

func findSubstring(s string, words []string) []int {

	lists := combination(words)
	var results []int
	for _, list := range  lists {
		if index := strings.Index(s, list); index != -1 {
			results = append(results, index)
		}
	}

	return results
}

func combination(words []string) []string {

	if len(words) <= 1 {
		return words
	}


	var lists []string
	for lix, word := range words {
		items := combination(append(words[:lix], words[lix+1:]...))
		for _, item := range items {
			lists = append(lists, word+item)
		}
	}
	return lists
}