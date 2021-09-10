package strStr

// 给你两个字符串 haystack 和 needle ，请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回  -1 。
// 思路: 遍历haystack, 如果遇到和字符needle第一个字符相同, 则挨个比较, 遇到不相同的则跳到下一个区间

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	lengthSrc, lengthDest := len(haystack), len(needle)
	for i := 0; i+lengthDest <= lengthSrc; i++ {
		if haystack[i] == needle[0] {
			k, j := i+1, 1
			for ; j < lengthDest; j++ {
				if haystack[k] != needle[j] {
					break
				}
				k++
			}
			if j == lengthDest {
				return i
			}
		}
	}

	return -1
}
