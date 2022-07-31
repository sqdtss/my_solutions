package main

import "fmt"

func main() {
	var N, M, K int
	fmt.Scan(&N, &M, &K)
	sz := make([]bool, M+2)
	for i := 0; i < M+2; i++ {
		sz[i] = true
	}
	for i := 0; i < K; i++ {
		var idx int
		fmt.Scan(&idx)
		sz[idx] = false
	}
	dp := make([][]int, M+2)
	for i := range dp {
		dp[i] = make([]int, N+1)
	}
	dp[0][N] = 1
	for i := 1; i <= M+1; i++ {
		if sz[i] == false {
			if i > 0 {
				for j := 1; j < N; j++ {
					dp[i][j] += dp[i-1][j+1]
				}
			}
			if i > 1 {
				for j := 1; j < N; j++ {
					dp[i][j] += dp[i-2][j+1]
				}
			}
			if i > 2 {
				for j := 1; j < N; j++ {
					dp[i][j] += dp[i-3][j+1]
				}
			}
		} else {
			if i > 0 {
				for j := 1; j < N+1; j++ {
					dp[i][j] += dp[i-1][j]
				}
			}
			if i > 1 {
				for j := 1; j < N+1; j++ {
					dp[i][j] += dp[i-2][j]
				}
			}
			if i > 2 {
				for j := 1; j < N+1; j++ {
					dp[i][j] += dp[i-3][j]
				}
			}
		}
	}
	ans := 0
	for i := 0; i < N+1; i++ {
		ans += dp[M+1][i]
	}
	fmt.Println(ans)
}
