package solution

import "math"

//132. 分割回文串 II
//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
//返回符合要求的 最少分割次数 。
//
//示例 1：
//输入：s = "aab"
//输出：1
//解释：只需一次分割就可将 s 分割成 ["aa","b"] 这样两个回文子串。
//
//示例 2：
//输入：s = "a"
//输出：0
//
//示例 3：
//输入：s = "ab"
//输出：1
//
//提示：
//1 <= s.length <= 2000
//s 仅由小写英文字母组成

func minCut(s string) int {
	n := len(s)
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = s[i] == s[j] && dp[i+1][j-1]
		}
	}
	f := make([]int, n)
	for i := range f {
		if dp[0][i] {
			continue
		} else {
			f[i] = math.MaxInt
			for j := 0; j < i; j++ {
				if dp[j+1][i] && f[j]+1 < f[i] {
					f[i] = f[j] + 1
				}
			}
		}
	}
	return f[n-1]
}
