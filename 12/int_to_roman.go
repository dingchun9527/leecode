package intToRoman

// 功能: 把整型数字变成罗马数字
// 思路: 罗马和10进制数无非也就是进制不同, 做个进制转换即可

func IntToRoman(num int) string {

	type item struct {
		unit  int
		label string
	}

	var mappingList = []item{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	result, carry := "", 0
	for _, item := range mappingList {
		carry = num / item.unit
		num = num % item.unit
		for i := 0; i < carry; i++ {
			result += item.label
		}
	}
	return result
}
