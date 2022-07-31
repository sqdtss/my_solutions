package main

import (
	"fmt"
)

type nodes struct {
	father   *nodes
	children []*nodes
	value    string
}

func main() {
	var in string
	fmt.Scan(&in)
	var strs []string
	start := 0
	for i := 0; i < len(in); i++ {
		if in[i] == '<' {
			start = i
		}
		if in[i] == '>' {
			strs = append(strs, in[start+1:i])
		}
	}
	head := &nodes{}
	cur := head
	for _, str := range strs {
		if str[0] != '/' && str[len(str)-1] != '/' {
			cur.children = append(cur.children, &nodes{father: cur})
			cur = cur.children[len(cur.children)-1]
			cur.value = str
		} else if str[len(str)-1] == '/' {
			cur.children = append(cur.children, &nodes{father: cur, value: str[:len(str)-1]})
		} else {
			cur = cur.father
		}
	}

	ans := ""
	var q []*nodes
	for _, child := range head.children {
		q = append(q, child)
	}
	for len(q) > 0 {
		top := q[0]
		q = q[1:]
		ans += top.value + " "
		for _, child := range top.children {
			q = append(q, child)
		}
	}
	if len(ans) > 0 {
		ans = ans[:len(ans)-1]
	}
	fmt.Println(ans)
}
