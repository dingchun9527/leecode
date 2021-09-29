package groupAnagrams

import (
	"sort"
)

// 功能: 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
// 思路: 异位词所包含的字母肯定是相同的, 将每个字符串字符进行排序


func groupAnagrams(strs []string) [][]string {

	mapping := make(map[string][]string)
	for _, str := range strs {
		key := sortString(str)
		if item, ok := mapping[key]; ok {
			item = append(item, str)
			mapping[key] = item
		} else {
			mapping[key] = []string{str}
		}
	}

	var result [][]string
	for _, item := range mapping {
		result = append(result, item)
	}
	return result
}


// 字符串排序
func sortString(str string) string {
	bytes := []byte(str)
	sort.Slice(bytes, func(i, j int) bool {return bytes[i] <bytes[j]})

	return string(bytes)
}