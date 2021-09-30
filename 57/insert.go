package insert

import merge "leecode/56"

// 功能： 插入区间
// 思路: 先找到可以插入的位置

func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	return merge.Merge(intervals)
}