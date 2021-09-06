package reverseInt

func Reverse(x int) int {
	result := 0
	for carry := 10; x != 0; x /= carry {
		count1 := x%carry
		result = result*carry + count1
	}

	if result > 2147483647 || result < -2147483648 {
		return 0
	}
	return result
}