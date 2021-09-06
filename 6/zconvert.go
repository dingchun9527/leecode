package zConvert

import "fmt"

// 功能: 给定一个字符串, 按照z字型进行排列
// 思路: 看图能发现规律是, 先下沉numRows个, 然后再回调
func Convert(s string, numRows int) string {
	contents := make([]string, numRows)

	for lix, liy, step := 0, 0, 1; lix < len(s); lix++ {
		contents[liy] = fmt.Sprintf("%s%c", contents[liy], s[lix])
		liy += step
		if liy >= numRows {
			step = -1
			liy = numRows - 2
			if numRows <= 1 {
				liy = numRows - 1
			}
		}

		if liy < 0 {
			step = 1

			liy = 1
			if numRows <= 1 {
				liy = 0
			}
		}
	}

	result := ""
	for lix := 0; lix < numRows; lix++ {
		result += contents[lix]
	}

	return result
}
