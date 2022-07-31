package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	type pair struct {
		index, num int
	}
	pairs := make([]pair, n)
	for i := 0; i < n; i++ {
		pairs[i].index = i
		fmt.Scan(&pairs[i].num)
	}
	sort.Slice(pairs, func(i, j int) bool {
		if pairs[i].num != pairs[j].num {
			return pairs[i].num < pairs[j].num
		}
		return pairs[i].index < pairs[j].index
	})
	ans := 0
	for {
		if pairs[0].num*len(pairs) < k {
			k -= pairs[0].num * len(pairs)
			for i := 1; i < len(pairs); i++ {
				pairs[i].num -= pairs[0].num
			}
			pairs = pairs[1:]
		} else {
			sort.Slice(pairs, func(i, j int) bool {
				return pairs[i].index < pairs[j].index
			})
			nowK := 0
			for {
				for i := 0; i < len(pairs); i++ {
					nowK++
					if nowK == k {
						ans = pairs[i].index + 1
						break
					}
				}
				if nowK == k {
					break
				}
			}
			break
		}
	}
	fmt.Println(ans)
}
