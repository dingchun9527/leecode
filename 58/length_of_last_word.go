package lengthOfLastWord

// 功能: 后一个单词的长度
// 思路: 遍历字符串, 遇到空格则重置长度

func lengthOfLastWord(s string) int {

	count, last := 0, 0
	for _, ch := range s {
		if ch == ' ' {
			if count != 0 {
				last = count
			}
			count = 0
			continue
		}
		count++
	}

	if count == 0 {
		count = last
	}
	return count
}