package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	var str string
	fmt.Scan(&str)
	sz := make([]int, n)
	for i, c := range str {
		if c == 'p' {
			sz[i] = 1
		} else if c == 'o' {
			sz[i] = 2
		} else if c == 'n' {
			sz[i] = 3
		} else {
			sz[i] = 4
		}
	}

	ans := 0
	nowI := 1
	for {
		flag := false
		for i := 0; i < len(sz); i++ {
			if sz[i] == nowI {
				nowI++
				sz[i] = 0
			}
		}
		if nowI == 4 {
			ans++
			nowI = 0
			flag = true
		}
		if !flag {
			break
		}
	}

	fmt.Println(ans)

}
