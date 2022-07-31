package solution

// 567. 字符串的排列
// 给你两个字符串 s1 和 s2 ，写一个函数来判断 s2 是否包含 s1 的排列。如果是，返回 true ；否则，返回 false 。
// 换句话说，s1 的排列之一是 s2 的 子串 。

// 示例 1：
// 输入：s1 = "ab" s2 = "eidbaooo"
// 输出：true
// 解释：s2 包含 s1 的排列之一 ("ba").

// 示例 2：
// 输入：s1= "ab" s2 = "eidboaoo"
// 输出：false

// 提示：
// 1 <= s1.length, s2.length <= 104
// s1 和 s2 仅包含小写字母

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}
	var s1Count, subS2Count [26]int
	for i := range s1 {
		s1Count[s1[i]-'a']++
		subS2Count[s2[i]-'a']++
	}
	if s1Count == subS2Count {
		return true
	}
	for i := len(s1); i < len(s2); i++ {
		subS2Count[s2[i-len(s1)]-'a']--
		subS2Count[s2[i]-'a']++
		if s1Count == subS2Count {
			return true
		}
	}
	return false
}
