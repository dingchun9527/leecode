package threeSum

import "sort"

// 功能: 给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组
// 思路: 先排序, 再取出第一个元素, 使用双指针扫描第二个和第三个, 因为第二个元素是递增, 第三个元素就一定是递减
func ThreeSum(nums []int) [][]int {

	// 输入检查
	if len(nums) < 3 {
		return [][]int{}
	}

	// 先排序
	sort.Sort(sort.IntSlice(nums))

	results := make([][]int, 0)
	length, sum := len(nums), 0
	// 遍历数组
	for i := 0; i < length; i++ {
		// 当前元素已大于0, 和后面的元素相加不可能等于0
		if nums[i] > 0 {
			break
		}

		// 跳过重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, length-1
		for left < right {
			sum = nums[i] + nums[left] + nums[right]
			if sum == 0 { // sum正好等于0, 剔除重复, 寻找下一个
				results = append(results, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				left++
				right--
			} else if sum > 0 { // sum大于0, 第三个元素需要减小
				right--
			} else { // sum小于0, 第二个元素需要加大
				left++
			}
		}
	}
	return results
}
