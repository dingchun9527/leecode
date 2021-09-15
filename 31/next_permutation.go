package nextPermutation

import "sort"

// 功能: 实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列（即，组合出下一个更大的整数）。
// 如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
// 必须 原地 修改，只允许使用额外常数空间。
// 思路: 首先, 把较大的数往左移动会使得整个数变大, 左移的位移越小这个数就越小, 那么我们的的结果是首先尽可能移动少的位移
// 在相同位移的情况下, 尽可能移动更小的数

func NextPermutation(nums []int)  {
	length := len(nums)
	if length <= 1 { return }

	var found bool
	var i, j int
	for i=length-2; i >= 0; i-- {
		// 从后面挑选一个较小的(位移尽可能的小), 让整个数变大
		for j = length-1; j>i; j-- {
			if nums[j] > nums[i] {
				nums[i], nums[j] = nums[j], nums[i]
				found = true
				break
			}
		}

		// 如果移动城管, 那么右边的部分要让他尽可能的小, 做一个从小到大排序
		if found {
			sort.Sort(sort.IntSlice(nums[i+1:]))
			break
		}
	}

	// 如果没找到更大的数则排序
	if !found {
		sort.Sort(sort.IntSlice(nums))
	}
}
