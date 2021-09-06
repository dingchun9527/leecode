package longestPalindrome

// 功能: 获取一个字符串的最长回文字串
// 思路: 动态规划, 长度是1的字符串一定是回文, 长度是2的只要两个字符相当一定是回文, 更长的转移方程为p(i,j) = p(i+1, j-1)&&(s[i] == s[j])

func LongestPalindrome(s string) string {
	length := len(s)

	// 长度小于等于1的一定是回文
	if length <= 1 {
		return s
	}

	// 申请数组保存状态, 长度为1的字符串一定是回文
	dp := make([][]bool, length)
	for lix := 0; lix < length; lix++ {
		dp[lix] = make([]bool, length)
	}

	for lix := 0; lix < length; lix++ {
		dp[lix][lix] = true
	}

	begin, maxLength := 0, 1
	// 枚举长度
	for L := 2; L <= length; L++ {
		// 遍历每个该长度的子串
		for lix := 0; lix < length; lix++ {
			liy := lix + L - 1
			if liy >= length {
				break
			}

			if s[lix] != s[liy] {
				dp[lix][liy] = false
			} else {
				if liy-lix < 3 {
					dp[lix][liy] = true
				} else {
					dp[lix][liy] = dp[lix+1][liy-1]
				}
			}

			if dp[lix][liy] && liy-lix+1 > maxLength {
				begin = lix
				maxLength = liy - lix + 1
			}
		}
	}
	return s[begin : begin+maxLength]
}
