package plusTwoNumber

// 功能: 求一个数组里的两个数之和等于目标值
// 思路: 遍历数组, 每个数计算出达到目标值需要匹配的数字, 将这个数字存到map中, 一旦遍历到某个数字发现他出现在匹配map中, 那么匹配成功ss

func TwoSum(nums []int, target int) []int {
	destMap := make(map[int]int)
	for lix, num := range nums {
		if liy, ok := destMap[num]; ok {
			return []int{liy, lix}
		}
		destMap[target-num] = lix
	}

	return []int{0, 0}
}
