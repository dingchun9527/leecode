package search

// 功能: 搜索旋转排序数组
// 要求: 时间复杂度O(log n)
// 思路: 虽然是经过翻转的, 但他还是局部有序的, 右边的比左边的部分要小, 先扫右边的部分, 如果比右边的第一个还要小, 那么肯定在右半部分, 否则在左半部分

func Search(nums []int, target int) int {
	length := len(nums)

	if target <= nums[length-1] {
		for i:=length-1; i>=0; i-- {
			if nums[i] == target {
				return i
			}
			if i > 0 && nums[i] < nums[i-1] || nums[i] < target{
				break
			}
		}
	} else {
		for i:=0; i<=length-1; i++ {
			if nums[i] == target {
				return i
			}
			if i < length-1 && nums[i] > nums[i+1] || nums[i] > target {
				break
			}
		}
	}
	return -1
}

