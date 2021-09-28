package permute

// 功能: 给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。
// 思路: 先取第一个数, 再取第二个 以此类推

func Permute(nums []int) [][]int {
	length := len(nums)

	var result [][]int
	if length == 1 {
		result = append(result, []int{nums[0]})
		return result
	}

	for i:=0; i<length; i++ {
		buffer := make([]int, length)
		copy(buffer, nums)
		list := Permute(append(buffer[:i],buffer[i+1:]...))
		for _, item := range list {
			result = append(result, append(item, nums[i]))
		}
	}

	return result
}

