package main

import (
	"fmt"
	"math"
)

func main() {
	var n, start, end, ttl int
	fmt.Scan(&n, &start, &end, &ttl)
	start--
	end--
	var graph [600][600]int
	for i := 0; i < n; i++ {
		var u, v, weight int
		fmt.Scan(&u, &v, &weight)
		graph[u-1][v-1] = weight
		graph[v-1][u-1] = weight
	}

	vis := make([]bool, n)

	maxTTL := 0
	ans := math.MaxInt64
	var dfs func(nowTTL, nowPos, nowAns int)
	dfs = func(nowTTL, nowPos, nowAns int) {
		if nowTTL < 0 {
			return
		}
		if nowPos == end {
			if nowAns < ans {
				ans = nowAns
				maxTTL = nowTTL
			} else if nowAns == ans {
				if nowTTL > maxTTL {
					maxTTL = nowTTL
				}
			}
			return
		}
		for i := 0; i < n; i++ {
			if graph[nowPos][i] > 0 && !vis[i] {
				vis[i] = true
				dfs(nowTTL-1, i, nowAns+graph[nowPos][i])
				vis[i] = false
			}
		}
	}
	vis[start] = true
	dfs(ttl, start, 0)
	if ans == math.MaxInt64 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans, maxTTL)
	}
}
