package search

// 功能: 搜索旋转排序数组
// 要求: 时间复杂度O(log n)
// 思路: 虽然是经过翻转的, 但他还是局部有序的, 右边的比左边的部分要小, 先扫右边的部分, 如果比右边的第一个还要小, 那么肯定在右半部分, 否则在左半部分
// 思路1: 二分查找

func Search(nums []int, target int) int {
	length := len(nums)

	if target <= nums[length-1] {
		for i := length - 1; i >= 0; i-- {
			if nums[i] == target {
				return i
			}
			if i > 0 && nums[i] < nums[i-1] || nums[i] < target {
				break
			}
		}
	} else {
		for i := 0; i <= length-1; i++ {
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

// 二分查找
// 通过二分一定可以找到一个有序的部分, 一个非有序的部分, 可以判断是否在有序的部分里来设置边界
func SearchNew(nums []int, target int) int {
	length := len(nums)
	if length <= 0 {
		return -1
	}

	if length == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}

	l, r := 0, length -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[l] <= nums[mid] {	// 左边有序
			if nums[l] <= target && target <= nums[mid] {	// 在有序数组里
				r = mid - 1
			} else {
				l = mid + 1
			}
		} else {	// 右边有序
			if nums[mid] <= target && target <= nums[r] {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
	}
	return -1
}
