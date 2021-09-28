package permuteUniquePack

// 功能: 给定一个可包含重复数字的序列 nums ，按任意顺序 返回所有不重复的全排列。
// 思路: 在46题的基础上做个去重

func PermuteUnique(nums []int) [][]int {
	length := len(nums)

	var result [][]int
	if length == 1 {
		result = append(result, []int{nums[0]})
		return result
	}

	mapping := make(map[int]struct{})
	for i:=0; i<length; i++ {
		if _, ok := mapping[nums[i]]; ok {
			continue
		}
		mapping[nums[i]] = struct{}{}
		buffer := make([]int, length)
		copy(buffer, nums)
		list := PermuteUnique(append(buffer[:i],buffer[i+1:]...))
		for _, item := range list {
			result = append(result, append(item, nums[i]))
		}
	}

	return result
}
