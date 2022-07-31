package main

import (
	"fmt"
	"math"
	"strconv"
	time2 "time"
)

func main() {
	var k int
	fmt.Scan(&k)
	mp := make(map[string][]int)
	mpI := make(map[int]string)
	c := 0
	for i := 0; i < k; i++ {
		var name, time string
		fmt.Scan(&name, &time)
		year := 2021
		month, _ := strconv.Atoi(time[5:7])
		day, _ := strconv.Atoi(time[8:10])
		if _, ok := mp[name]; !ok {
			mpI[c] = name
			c++
		}
		mp[name] = append(mp[name], time2.Date(year, time2.Month(month), day, 0, 0, 0, 0, time2.UTC).Day())
	}
	var n int
	fmt.Scan(&n)
	if len(mp) < n {
		fmt.Println(-1)
	} else {
		ans := math.MaxInt64
		var dfs func(index, nowN, nowMinTime, nowMaxTime int)
		dfs = func(index, nowN, nowMinTime, nowMaxTime int) {
			if index >= c {
				return
			}
			if nowN == n {
				if nowMaxTime-nowMinTime+1 < ans {
					ans = nowMaxTime - nowMinTime + 1
				}
				return
			}

			for i := 0; i < len(mp[mpI[index]]); i++ {
				tmpNowMinTime, tmpNowMaxTime := nowMinTime, nowMaxTime
				if mp[mpI[index]][i] < tmpNowMinTime {
					tmpNowMinTime = mp[mpI[index]][i]
				}
				if mp[mpI[index]][i] > tmpNowMaxTime {
					tmpNowMaxTime = mp[mpI[index]][i]
				}
				dfs(index+1, nowN+1, tmpNowMinTime, tmpNowMaxTime)
			}
			dfs(index+1, nowN, nowMinTime, nowMaxTime)
		}
		dfs(0, 0, math.MaxInt64, math.MinInt64)
		fmt.Println(ans)
	}

}
