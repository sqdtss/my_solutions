package main

import (
	"fmt"
	"strings"
)

//面试题 17.07. 婴儿名字
//每年，政府都会公布一万个最常见的婴儿名字和它们出现的频率，也就是同名婴儿的数量。有些名字有多种拼法，
//例如，John 和 Jon 本质上是相同的名字，但被当成了两个名字公布出来。给定两个列表，一个是名字及对应的频率，另一个是本质相同的名字对。
//设计一个算法打印出每个真实名字的实际频率。注意，如果 John 和 Jon 是相同的，并且 Jon 和 Johnny 相同，则 John 与 Johnny 也相同，即它们有传递和对称性。
//在结果列表中，选择 字典序最小 的名字作为真实名字。
//
//示例：
//输入：names = ["John(15)","Jon(12)","Chris(13)","Kris(4)","Christopher(19)"],
//	 synonyms = ["(Jon,John)","(John,Johnny)","(Chris,Kris)","(Chris,Christopher)"]
//输出：["John(27)","Chris(36)"]
//
//提示：
//names.length <= 100000

func trulyMostPopular(names []string, synonyms []string) (ans []string) {
	parent := make(map[string]string)
	count := make(map[string]int)
	for _, name := range names {
		n, c := getNameAndCount(name)
		parent[n] = n
		count[n] = c
	}
	for _, synonym := range synonyms {
		A, B := getAAndB(synonym)
		parent1, parent2 := find(A, parent), find(B, parent)
		if parent1 == "" {
			parent[parent1] = parent2
			continue
		}
		if parent2 == "" {
			parent[parent2] = parent1
			continue
		}

		if parent1 > parent2 {
			parent1, parent2 = parent2, parent1
		}
		//将parent2 并入 parent1
		if parent1 != parent2 {
			count[parent1] += count[parent2]
			parent[parent2] = parent1
			delete(count, parent2)
		}
	}

	for name, cnt := range count {
		s := fmt.Sprintf("%s(%d)", name, cnt)
		ans = append(ans, s)
	}
	return
}

// 找到名字对应的祖先名字， 并使用路径压缩进行优化
func find(s string, parent map[string]string) string {
	if parent[s] != s {
		parent[s] = find(parent[s], parent)
	}
	return parent[s]
}

func getNameAndCount(s string) (name string, count int) {
	flag := false
	for _, c := range s {
		if c == '(' {
			flag = true
		} else if c == ')' {
			break
		} else {
			if flag {
				count = count*10 + int(c-'0')
			} else {
				name += string(c)
			}
		}
	}
	return
}

func getAAndB(s string) (string, string) {
	tmp := strings.Split(s[1:len(s)-1], ",")
	return tmp[0], tmp[1]
}
