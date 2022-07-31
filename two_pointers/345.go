package solution

import "strings"

// 345. 反转字符串中的元音字母
// 给你一个字符串 s ，仅反转字符串中的所有元音字母，并返回结果字符串。
// 元音字母包括 'a'、'e'、'i'、'o'、'u'，且可能以大小写两种形式出现。

// 示例 1：
// 输入：s = "hello"
// 输出："holle"

// 示例 2：
// 输入：s = "leetcode"
// 输出："leotcede"

// 提示：
// 1 <= s.length <= 3 * 105
// s 由 可打印的 ASCII 字符组成

func reverseVowels(s string) string {
	tmp := []byte(s)
	target := "aeiouAEIOU"
	i, j := 0, len(tmp)-1
	for i < j {
		for i < j && !strings.Contains(target, string(tmp[i])) {
			i++
		}
		for i < j && !strings.Contains(target, string(tmp[j])) {
			j--
		}
		if i < j {
			tmp[i], tmp[j] = tmp[j], tmp[i]
			i++
			j--
		}
	}
	return string(tmp)
}
