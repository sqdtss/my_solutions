package solution

//131. 分割回文串
//给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//回文串 是正着读和反着读都一样的字符串。
//
//示例 1：
//输入：s = "aab"
//输出：[["a","a","b"],["aa","b"]]
//
//示例 2：
//输入：s = "a"
//输出：[["a"]]
//
//提示：
//1 <= s.length <= 16
//s 仅由小写英文字母组成

func partition(s string) (res [][]string) {
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
	split := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			tmp := make([]string, len(split))
			copy(tmp, split)
			res = append(res, tmp)
			return
		}
		for j := i; j < n; j++ {
			if dp[i][j] {
				split = append(split, s[i:j+1])
				dfs(j + 1)
				split = split[:len(split)-1]
			}
		}
	}
	dfs(0)
	return
}
