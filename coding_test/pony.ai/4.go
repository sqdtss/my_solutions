package main

import (
	"fmt"
	"math"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}
	for i := 0; i+m <= n; i++ {
		min := math.MaxInt64
		for j := 0; j < m; j++ {
			if nums[i+j] < min {
				min = nums[i+j]
			}
		}
		for j := 0; j < m; j++ {
			nums[i+j] -= min
		}
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)
}
