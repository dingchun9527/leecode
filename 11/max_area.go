package maxArea

// 功能: 盛最多水的容器
// 思路: 双指针, 移动较小的那一个, 如果遇到更大的则计算一下

func maxArea(height []int) int {
	max := 0
	for i, j := 0, len(height)-1; i < j; {
		curr := min(height[i], height[j]) * (j - i)
		if curr > max {
			max = curr
		}

		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
