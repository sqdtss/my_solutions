package solution

/* solution
472. 连接词
给你一个 不含重复 单词的字符串数组 words ，请你找出并返回 words 中的所有 连接词 。

连接词 定义为：一个完全由给定数组中的至少两个较短单词组成的字符串。

示例 1：
输入：words = ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]
输出：["catsdogcats","dogcatsdog","ratcatdogcat"]
解释："catsdogcats" 由 "cats", "dog" 和 "cats" 组成;
     "dogcatsdog" 由 "dog", "cats" 和 "dog" 组成;
     "ratcatdogcat" 由 "rat", "cat", "dog" 和 "cat" 组成。

示例 2：
输入：words = ["cat","dog","catdog"]
输出：["catdog"]

*/
func findAllConcatenatedWordsInADict(words []string) (ans []string) {
	mp := make(map[int64]int)
	P, OFFSET := 131, 128
	n := len(words)

	check := func(s string) bool {
		n := len(s)
		f := make([]int, n+1)
		for i := range f {
			f[i] = -1
		}
		f[0] = 0
		for i := 0; i <= n; i++ {
			if f[i] == -1 {
				continue
			}
			var cur int64
			for j := i + 1; j <= n; j++ {
				cur = cur*int64(P) + int64(s[j-1]-'a') + int64(OFFSET)
				if mp[cur] != 0 {
					f[j] = max(f[j], f[i]+1)
				}
			}
			if f[n] > 1 {
				return true
			}
		}
		return false
	}

	for i := 0; i < n; i++ {
		var hash int64
		for _, c := range words[i] {
			hash = hash*int64(P) + int64(c-'a') + int64(OFFSET)
		}
		mp[hash]++
	}
	for _, word := range words {
		if check(word) {
			ans = append(ans, word)
		}
	}
	return
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// 暴力做法

// func findAllConcatenatedWordsInADict(words []string) (ans []string) {

// 	mp := make(map[string]int)

// 	check := func(s string) bool {
// 		n := len(s)
// 		f := make([]int, n+1)
// 		for i := range f {
// 			f[i] = -1
// 		}
// 		f[0] = 0
// 		for i := 0; i <= n; i++ {
// 			if f[i] == -1 {
// 				continue
// 			}
// 			for j := i + 1; j <= n; j++ {
// 				if mp[s[i:j]] != 0 {
// 					f[j] = max(f[j], f[i]+1)
// 				}
// 			}
// 			if f[n] > 1 {
// 				return true
// 			}
// 		}
// 		return false
// 	}

// 	for _, word := range words {
// 		mp[word]++
// 	}
// 	for _, word := range words {
// 		if check(word) {
// 			ans = append(ans, word)
// 		}
// 	}
// 	return
// }
