package combinationSum

import "sort"

// 功能: 给定一个无重复元素的正整数数组 candidates 和一个正整数 target ，找出 candidates 中所有可以使数字和为目标数 target 的唯一组合。
// 思路: 先排序

func combinationSum(candidates []int, target int) [][]int {

	// 先排序
	sort.Sort(sort.IntSlice(candidates))

	// 从第一个数开始尝试
	for i:=0; i<len(candidates); i++ {

		// 如果一个数已经大于target, 那后面的数相加已经不可能等于target了
		if candidates[i] > target {
			break
		}

		match := combinationSum(candidates[i-1])
	}
}
