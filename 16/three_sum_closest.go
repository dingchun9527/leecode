package threeSumClosest

import (
	"math"
	"sort"
)

// 功能: 给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。
// 思路: 先给数组排个序, 然后遍历该数组, 通过左右移动两个元素得到最接近的结果

func ThreeSumClosest(nums []int, target int) int {
	// 输入检查
	if len(nums) < 3 {
		return 0
	}

	// 排序
	sort.Sort(sort.IntSlice(nums))

	// 遍历数组
	dest, length := nums[0]+nums[1]+nums[2], len(nums)
	for i := 0; i < length; i++ {

		// 如果当前元素已经大于target&大于0, 那么再和后面的元素相加会更大
		if nums[i] > target && nums[i] > 0 {
			break
		}

		// 剔除重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, length-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target { // 如果相等, 那么结束了
				return sum
			} else if sum > target { // sum较大, 那么第三个元素过大, 适当调小
				if math.Abs(float64(sum-target)) < math.Abs(float64(dest-target)) {
					dest = sum
				}
				right--
			} else {
				if math.Abs(float64(sum-target)) < math.Abs(float64(dest-target)) {
					dest = sum
				}
				left++
			}
		}
	}

	return dest
}
