package threeSum

import "sort"

// 功能: 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组
// 思路: 先排序, 再取出第一个元素, 使用双指针扫描第二个和第三个, 因为第二个元素是递增, 第三个元素就一定是递减
func ThreeSum(nums []int) [][]int {

	// 输入检查
	if len(nums) < 3 {
		return [][]int{}
	}

	// 排序
	sort.Sort(sort.IntSlice(nums))

	results, sum, length := make([][]int, 0), 0, len(nums)

	// 遍历数组
	for i := 0; i < length; i++ {
		// 如果当前元素已经大于0, 不可能和更大的元素相加得到0
		if nums[i] > 0 {
			break
		}

		// 剔除重复的情况
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// 双指针获取后两个元素
		left, right := i+1, length-1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if sum == 0 { // sum刚好等于0, 需要往后移动第二个元素(变大), 往左移动第三个元素(变小)
				results = append(results, []int{nums[i], nums[left], nums[right]}) // 加入结果集
				if left < right && nums[right] == nums[right-1] {                  // 剔除重复元素
					right--
				}

				for left < right && nums[left] == nums[left+1] {
					left++
				}

				left++  // 第二个元素右移
				right-- // 第三个元素左移
			} else if sum > 0 { // 第二个元素过小, 有移
				right--
			} else { // 第三个元素过大, 左移
				left++
			}
		}
	}
	return results
}
