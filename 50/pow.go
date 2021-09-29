package pow

// 思想: 通过快速幂的方式, 快速迫近结果, 而不是每次只乘以一次

func myPow(x float64, n int) float64 {
	if n >= 0 {
		return quickMul(x, n)
	}

	return 1.0/quickMul(x, -n)
}

func quickMul(x float64, n int) float64 {
	if n == 0 {
		return 1
	}

	y := quickMul(x, n/2)
	if n%2 == 0 {
		return y * y
	}
	return y * y * x
}