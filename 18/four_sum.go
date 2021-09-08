package fourSum

import "sort"

// 功能: 给你一个由 n 个整数组成的数组 nums ，和一个目标值 target 。请你找出并返回满足下述全部条件且不重复的四元组 [nums[a], nums[b],
// 思路:

func FourSum(nums []int, target int) [][]int {

	// 长度校验
	if len(nums) < 4 {
		return [][]int{}
	}

	// 数组排序
	sort.Sort(sort.IntSlice(nums))

	// 遍历数组
	results, length := make([][]int, 0), len(nums)
	for i := 0; i < length-3; i++ {
		// 连续4个数已大于target, 那么后续的数家起来不可能等于target
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			break
		}

		// 当前数字过小, 跳过本次循环
		if nums[i]+nums[length-3]+nums[length-2]+nums[length-1] < target {
			continue
		}

		// 剔除重复的元素
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 遍历第二个元素
		for j := i + 1; j < length-2; j++ {
			// 剔除重复的元素
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			// 连续四个数已大于target, 则往后的数更大, 跳出本次循环
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}

			// 和最后两个数组合还小于target, 再往后加别的数也不可能等于target
			if nums[i]+nums[j]+nums[length-2]+nums[length-1] < target {
				continue
			}

			// 使用双指针移动后面的两个数
			left, right := j+1, length-1
			for left < right {
				sum := nums[i] + nums[j] + nums[left] + nums[right]
				if sum == target { // 刚好匹配, left右移动, right左移
					results = append(results, []int{nums[i], nums[j], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] { // 剔除重复元素
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}
					left++
					right--
				} else if sum > target { // sum 过大, 需要right 左移
					right--
				} else {
					left++
				}
			}
		}
	}
	return results
}
