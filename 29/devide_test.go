package devide

// 功能:  给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
// 返回被除数 dividend 除以除数 divisor 得到的商
// 思路: 使用减法来实现

func divide(dividend int, divisor int) int {

	// 被除数检查
	if divisor == 0 {
		panic(divisor)
	}

	// 除数检查
	if dividend == 0 {
		return 0
	}

	flag := 1
	if dividend > 0 && divisor < 0 || dividend <0 && divisor > 0 {
		flag = -1
	}

	// 取两个数的绝对值
	if dividend < 0 {
		dividend *= -1
	}

	if divisor < 0 {
		divisor *= -1
	}

	// 两数相减, 直至被除数小于除数
	result, count := 0, 1
	base := divisor
	for ;dividend >= divisor; dividend, divisor= dividend-divisor, divisor + divisor {
		result += count
		count += count
		//     fmt.Printf("dividend:%d divisor:%d\n", dividend, divisor)
	}

	for dividend < divisor && count > 1 {
		divisor -= base
		count--
		//      fmt.Printf("111 dividend:%d divisor:%d\n", dividend, divisor)

	}
	//   fmt.Printf("count:%d\n", count)

	if count >= 1 {
		for i:=0; i<count; i++ {
			for ;dividend >= divisor ; dividend= dividend-divisor {
				//                       fmt.Printf("222 dividend:%d divisor:%d\n", dividend, divisor)

				result += count
			}
			divisor -= base
			count--
		}
	}

	result *= flag
	if result > 2147483647 {
		result = 2147483647
	}

	if result < -2147483648 {
		result = -2147483648
	}

	return result
}

// TODO 需要有更好的优化方案, 提高实践效率

