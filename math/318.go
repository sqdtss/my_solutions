package solution

// 318. 最大单词长度乘积
// 给你一个字符串数组 words ，找出并返回 length(words[i]) * length(words[j]) 的最大值，并且这两个单词不含有公共字母。如果不存在这样的两个单词，返回 0 。

// 示例 1：
// 输入：words = ["abcw","baz","foo","bar","xtfn","abcdef"]
// 输出：16
// 解释：这两个单词为 "abcw", "xtfn"。

// 示例 2：
// 输入：words = ["a","ab","abc","d","cd","bcd","abcd"]
// 输出：4
// 解释：这两个单词为 "ab", "cd"。

// 示例 3：
// 输入：words = ["a","aa","aaa","aaaa"]
// 输出：0
// 解释：不存在这样的两个单词。

// 提示：
// 2 <= words.length <= 1000
// 1 <= words[i].length <= 1000
// words[i] 仅包含小写字母

func maxProduct(words []string) (ans int) {
	wordCode := make([]int, len(words))
	for i, word := range words {
		for _, c := range word {
			wordCode[i] |= (1 << (c - 'a'))
		}
	}
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if wordCode[i]&wordCode[j] == 0 && len(words[i])*len(words[j]) > ans {
				ans = len(words[i]) * len(words[j])
			}
		}
	}
	return
}
