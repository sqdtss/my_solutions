package main

func getLongestPalindrome(A string) int {
	ans := 1
	n := len(A)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = true
	}
	for i := 0; i < n-1; i++ {
		if A[i] == A[i+1] {
			dp[i][i+1] = true
			ans = 2
		}
	}
	for k := 2; k < n; k++ {
		for i := 0; i < n-k; i++ {
			if A[i] == A[i+k] && dp[i+1][i+k-1] {
				ans = k + 1
				dp[i][i+k] = true
			}
		}
	}
	return ans
}
