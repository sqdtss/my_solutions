package solution

import "strings"

//面试题 01.03. URL化
//URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，
//并且知道字符串的“真实”长度。
//示例 1：
//输入："Mr John Smith    ", 13
//输出："Mr%20John%20Smith"
//
//示例 2：
//输入："               ", 5
//输出："%20%20%20%20%20"
//
//提示：
//字符串长度在 [0, 500000] 范围内。

func replaceSpaces(S string, length int) string {
	ans := strings.Replace(S[:length], " ", "%20", -1)
	return ans
}
