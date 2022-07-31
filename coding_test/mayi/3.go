package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	graph := make([][]int, n)
	for i := range graph {
		graph[i] = make([]int, n)
	}
	vis := make([]bool, n*n)
	ok := false
	var ans [][]int
	var dfs func(nowIndex int)
	dfs = func(nowIndex int) {
		if ok {
			return
		}
		if nowIndex == n*n {
			ok = true
			ans = make([][]int, n)
			for i := range ans {
				ans[i] = make([]int, n)
				for j := 0; j < n; j++ {
					ans[i][j] = graph[i][j]
				}
			}
			return
		}
		x, y := nowIndex/n, nowIndex%n
		for i := 0; i < n*n; i++ {
			if !vis[i] {
				if x > 0 && y > 0 {
					tmp := graph[x-1][y] + graph[x][y-1] + graph[x-1][y-1] + i + 1
					if tmp%2 == 0 {
						continue
					}
				}

				vis[i] = true
				graph[x][y] = i + 1
				dfs(nowIndex + 1)
				vis[i] = false
			}
		}
	}
	dfs(0)
	if ok {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				fmt.Print(ans[i][j])
				if j == n-1 {
					fmt.Println()
				} else {
					fmt.Print(" ")
				}
			}
		}
	} else {
		fmt.Println(-1)
	}
}
