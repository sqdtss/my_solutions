package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var graph [2000 + 100][2000 + 100]byte
	for i := 0; i < n; i++ {
		var tmp string
		fmt.Scan(&tmp)
		for j := 0; j < m; j++ {
			graph[i][j] = tmp[j]
		}
	}
	ans := 0
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}
	var dfs func(x, y, nowAns int)
	dfs = func(x, y, nowAns int) {
		if x < 0 || x >= n || y < 0 || y >= m || vis[x][y] {
			if nowAns > ans {
				ans = nowAns
			}
			return
		}
		vis[x][y] = true
		if graph[x][y] == '^' {
			dfs(x-1, y, nowAns+1)
		} else if graph[x][y] == '<' {
			dfs(x, y-1, nowAns+1)
		} else if graph[x][y] == '>' {
			dfs(x, y+1, nowAns+1)
		} else {
			dfs(x+1, y, nowAns+1)
		}
		vis[x][y] = false
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			dfs(i, j, 0)
		}
	}
	fmt.Println(ans)
}
