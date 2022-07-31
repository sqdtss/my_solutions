package solution

//迭代快速幂
func pow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return x
	}
	var isNeg bool
	if n < 0 {
		isNeg = true
		n = -n
	}
	ans, base := 1.0, x
	for n > 0 {
		if n&1 == 1 {
			ans *= base
		}
		base *= base
		n = n >> 1
	}
	if isNeg {
		return 1.0 / ans
	}
	return ans
}
