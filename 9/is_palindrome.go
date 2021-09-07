package isPalindrome

// 功能：判断一个整数是否是回文数
// 思想: 将一个整数放进一个slince, 然后判断是否是回文
func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	// 整数拆分成数组
	var array []int
	for carry := x; carry != 0; carry /= 10 {
		array = append(array, carry%10)
	}

	// 判断数组是否为回文
	return isPalindromeArray(array)
}

func isPalindromeArray(array []int) bool {
	if len(array) <= 1 {
		return true
	}

	for i, j := 0, len(array)-1; i < j; i, j = i+1, j-1 {
		if array[i] != array[j] {
			return false
		}
	}

	return true
}
