package solution

//97. 交错字符串
//给定三个字符串 s1、s2、s3，请你帮忙验证 s3 是否是由 s1 和 s2 交错 组成的。
//两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
//s = s1 + s2 + ... + sn
//t = t1 + t2 + ... + tm
//|n - m| <= 1
//交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
//注意：a + b 意味着字符串 a 和 b 连接。
//
//示例 1：
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbcbcac"
//输出：true
//
//示例 2：
//输入：s1 = "aabcc", s2 = "dbbca", s3 = "aadbbbaccc"
//输出：false
//
//示例 3：
//输入：s1 = "", s2 = "", s3 = ""
//输出：true
//
//提示：
//0 <= s1.length, s2.length <= 100
//0 <= s3.length <= 200
//s1、s2、和 s3 都由小写英文字母组成

func isInterleave(s1 string, s2 string, s3 string) bool {
	lenS1, lenS2 := len(s1), len(s2)
	if lenS1+lenS2 != len(s3) {
		return false
	}
	dp := make([][]bool, lenS1+1)
	for i := range dp {
		dp[i] = make([]bool, lenS2+1)
	}
	dp[0][0] = true
	for i := 1; i <= lenS1 && s1[i-1] == s3[i-1]; i++ {
		dp[i][0] = true
	}
	for i := 1; i <= lenS2 && s2[i-1] == s3[i-1]; i++ {
		dp[0][i] = true
	}
	for i := 1; i <= lenS1; i++ {
		for j := 1; j <= lenS2; j++ {
			dp[i][j] = (dp[i][j-1] && s2[j-1] == s3[i+j-1]) || (dp[i-1][j] && s1[i-1] == s3[i+j-1])
		}
	}
	return dp[lenS1][lenS2]
}
