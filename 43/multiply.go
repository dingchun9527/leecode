package multiply

import "fmt"

// 功能: 字符串相乘
// 思路: 把被乘数每一位挨个和乘数相乘, 结果相加

func multiply(num1 string, num2 string) string {

	// 0乘以任何数得0
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	// 遍历被乘数
	var result, unit, unit1 int32
	unit1 = 1
	for i:=len(num1)-1; i>=0; i-- {

		unit = 1
		for j:=len(num2)-1; j>=0; j-- {
			tmp := (int32(num1[i]) - '0') * (int32(num2[j]) - '0') * unit * unit1
			result += tmp
			unit *= 10
		}
		unit1 *= 10
	}
	return fmt.Sprintf("%d", result)
}