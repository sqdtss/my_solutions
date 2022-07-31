package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	bs, _, _ := rd.ReadLine()
	n, _ := strconv.Atoi(string(bs))
	var graph [100 + 10][100 + 10]bool
	mp := make(map[string]int)
	var in [100 + 10]int
	index := 0
	for i := 0; i < n; i++ {
		bs, _, _ = rd.ReadLine()
		sz := strings.Split(string(bs), " ")
		for j := 1; j < len(sz); j++ {
			var p, q int
			if _, ok := mp[sz[j-1]]; !ok {
				mp[sz[j-1]] = index
				p = index
				index++
			} else {
				p = mp[sz[j-1]]
			}
			if _, ok := mp[sz[j]]; !ok {
				mp[sz[j]] = index
				q = index
				index++
			} else {
				q = mp[sz[j]]
			}
			graph[p][q] = true
			in[q]++
		}
	}
	idx := 0
	for ; idx < len(in); idx++ {
		if in[idx] == 0 {
			break
		}
	}
	for k, v := range mp {
		if v == idx {
			fmt.Println(k)
		}
	}
	bs, _, _ = rd.ReadLine()
	level, _ := strconv.Atoi(string(bs))
	ans := 0
	var dfs func(nowIndex, nowLevel int)
	dfs = func(nowIndex, nowLevel int) {
		if nowLevel == level {
			ans++
			return
		}
		for idx, v := range graph[nowIndex] {
			if v == true {
				dfs(idx, nowLevel+1)
			}
		}
	}
	dfs(idx, 1)
	fmt.Println(ans)
}
