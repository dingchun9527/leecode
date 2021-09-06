package myAtoi

// 功能: 实现一个Atoi
// 思路: 挨个读取, 在转换
func MyAtoi(s string) int {
	var result float64
	flag, isNumber :=' ', false
	for _, ch := range s {

		// 去除前导空格
		if (ch == ' ') && !isNumber {
			continue
		}

		// 去除前导0
		if (ch == '0') && result == 0 {
			isNumber = true
			continue
		}

		// 符号判断
		if ch == '-' || ch == '+'{
			if isNumber && result == 0.0 {
				return 0
			}
			if flag != ' ' {
				break
			}

			if result != 0.0 {
				continue
			}
			flag = ch
			isNumber = true
			continue
		}

		// 去除非法字符
		if ch < '0' || ch > '9' {
			break
		}

		result = result *10 + float64(ch) - '0'
		isNumber = true
	}

	if flag == '-' {
		result *= -1
	}

	if result > 2147483647 {
		return 2147483647
	}

	if  result < -2147483648 {
		return -2147483648
	}

	return int(result)
}
