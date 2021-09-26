package searchInsert

// 功能: 搜索插入位置
// 要求: 时间复杂度为O(log n)


func SearchInsert(nums []int, target int) int {
	length := len(nums)

	l, r, mid, ans := 0, length - 1, 0, length

	for l <= r {
		mid = (l + r) / 2
		if target <= nums[mid] {
			ans = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}
