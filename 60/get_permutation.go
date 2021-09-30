package getPermutation

import "fmt"

// 功能: 排列序列
// 思路: 使用递归的范式进行排列, 返回第K个值

func getPermutation(n int, k int) string {
	array := make([]int, n)
	for i:=0; i<n; i++ {
		array[i] = i+1
	}

	list := permutation(array)

	item := list[k-1]
	str := ""
	for _, ch := range item {
		str = fmt.Sprintf("%s%d", str, ch)
	}

	return str
}

func permutation(array []int)[][]int {
	if len(array) == 1 {
		return [][]int{array}
	}

	var result [][]int
	for i:=0; i<len(array); i++ {
		buffer := make([]int, len(array))
		copy(buffer, array)
		list := permutation(append(buffer[:i],buffer[i+1:]...))

		for _, item := range list {
			var buffer []int
			buffer = append(buffer, array[i])
			buffer = append(buffer, item...)
			result = append(result, buffer)
		}
	}


	return result
}
