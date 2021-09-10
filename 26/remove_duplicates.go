package removeDuplicates

// 功能:  删除有序数组中的重复项
// 要求: 空间复杂度为O(1)
// 思路: 拿一个数字放到不重复的位置, 遇到重复的直接跳过, 遇到不重复的积蓄往不重复的位置后面放

func RemoveDuplicates(nums []int) int {

	// 异常判断
	if len(nums) <= 0 {
		return 0
	}

	// 去除重复元素
	noDupCount := 0
	for i:=0; i<len(nums);i++ {
		if nums[i] == nums[noDupCount] {
			continue
		}
		noDupCount++
		nums[noDupCount] = nums[i]
	}

	return noDupCount+1
}