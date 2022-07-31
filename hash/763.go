package solution

// 763. 划分字母区间
// 字符串 S 由小写字母组成。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。返回一个表示每个字符串片段的长度的列表。

// 示例：
// 输入：S = "ababcbacadefegdehijhklij"
// 输出：[9,7,8]
// 解释：
// 划分结果为 "ababcbaca", "defegde", "hijhklij"。
// 每个字母最多出现在一个片段中。
// 像 "ababcbacadefegde", "hijhklij" 的划分是错误的，因为划分的片段数较少。

// 提示：
// S的长度在[1, 500]之间。
// S只包含小写字母 'a' 到 'z'。

func partitionLabels(s string) (ans []int) {
	var eachCount [26]int
	for _, c := range s {
		eachCount[c-'a']++
	}
	mp := make(map[byte]bool)
	nowLen := 0
	for _, c := range s {
		mp[byte(c)] = true
		nowLen++
		eachCount[c-'a']--
		flag := true
		for k := range mp {
			if mp[k] && eachCount[k-'a'] != 0 {
				flag = false
				break
			}
		}
		if flag {
			ans = append(ans, nowLen)
			nowLen = 0
		}
	}
	return
}
