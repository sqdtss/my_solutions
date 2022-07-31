package main

import "fmt"

func main() {
	var t int
	for fmt.Scan(&t); t > 0; t-- {
		var n int
		fmt.Scan(&n)
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&nums[i])
		}
		ans := 1
		dp := make([][]bool, n)
		for i := range dp {
			dp[i] = make([]bool, n)
		}
		for i := 0; i < n; i++ {
			dp[i][i] = true
		}
		for i := 0; i+1 < n; i++ {
			if nums[i] == nums[i+1] {
				dp[i][i+1] = true
				ans = 2
			}
		}
		for k := 2; k < n; k++ {
			for i := 0; i+k < n; i++ {
				if nums[i] == nums[i+k] && nums[i] <= nums[i+1] {
					dp[i][i+k] = dp[i+1][i+k-1]
				}
				if dp[i][i+k] {
					ans = k + 1
				}
			}
		}
		fmt.Println(ans)
	}
}
