package my_template

import "math"

func Max(param ...int) int {
	max := math.MinInt
	for _, i := range param {
		if i > max {
			max = i
		}
	}
	return max
}

func Min(param ...int) int {
	min := math.MaxInt
	for _, i := range param {
		if i < min {
			min = i
		}
	}
	return min
}

func Abs(param int) int {
	if param < 0 {
		param = -param
	}
	return param
}

func Gcd(a, b int) int {
	if a%b == 0 {
		return b
	}
	return Gcd(b, a%b)
}
