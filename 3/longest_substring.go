package longestSubstring

import "fmt"

// 功能: 求不重复的最长子串的长度
// 思路: 滑动窗口实现
func LengthOfLongestSubstring(s string) int {
	sub := ""
	longest := 0
	for _, ch := range s {
		for lix, ch1 := range sub {
			if ch == ch1 {
				sub = sub[lix+1:]
				break
			}
		}
		sub = fmt.Sprintf("%s%c", sub, ch)
		size := len(sub)
		if size > longest {
			longest = size
		}
	}

	return longest
}

// TODO 双指针
