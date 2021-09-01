package plusTwoNumber

func TwoSum(nums []int, target int) []int {
	destMap := make(map[int]int)
	for lix, num := range nums {
		if liy, ok := destMap[num]; ok {
			return []int{liy, lix}
		}
		destMap[target -num ] = lix
	}

	return []int{0, 0}
}