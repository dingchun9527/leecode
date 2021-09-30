package merge

import "sort"

// 功能: 合并区间
// 思路: 使用第一个元素进行排序, 可合并的区间一定是连续的

func merge(intervals [][]int) [][]int {

	if len(intervals) <= 1 {
		return intervals
	}

	// 先进行排序
	sort.Slice(intervals, func(i, j int)bool{return intervals[i][0] < intervals[j][0]})

	// 遍历排序后的数组
	var result [][]int
	curr := intervals[0]
	for i:=1; i<len(intervals); i++ {
		if intervals[i][0] <= curr[1] {
			right := curr[1]
			if intervals[i][1] > right {
				right = intervals[i][1]
			}

			curr =  []int{curr[0], right}
		} else {
			result = append(result, curr)
			curr = intervals[i]
		}
		if i == len(intervals) - 1{
			result = append(result, curr)
		}
	}
	return result
}