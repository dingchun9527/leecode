package removeElement

// 功能: 给你一个数组 nums 和一个值 val，你需要 原地 移除所有数值等于 val 的元素，并返回移除后数组的新长度。
// 思路: 遇到不等于val的元素自从0开始覆盖, 最后把多余的数据剪掉

func RemoveElement(nums []int, val int) int {
	count := 0
	for _, num := range nums {
		if num != val {
			nums[count] = num
			count++
		}
	}
	nums = nums[:count]
	return count
}
