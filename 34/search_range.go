package searchRange

// 功能: 在排序数组中查找元素的第一个和最后一个位置
// 要求: 时间复杂度为O(log n)

func searchRange(nums []int, target int) []int {
	length := len(nums)
	if length <= 0 { return []int{-1, -1}}

	if length == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}

	// 二分查找
	l, r, mid := 0, length - 1, 0
	result := []int{-1, -1}

	for l <= r {
		mid = (l + r) / 2
		if nums[mid] == target {
			result[0], result[1] = mid, mid
			for i:=mid-1; i>=0 && nums[i] == target; i-- {
				result[0] = i
			}
			for i:=mid+1; i<=length-1 && nums[i] == target; i++ {
				result[1] = i
			}
		}

		if nums[l] <= target && target <= nums[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return result
}
