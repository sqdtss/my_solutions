package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	rd := bufio.NewReader(os.Stdin)
	bs, _, _ := rd.ReadLine()
	str := string(bs)
	weightsStr := strings.Split(str, " ")
	var weights []int
	for _, weightStr := range weightsStr {
		weight, _ := strconv.Atoi(weightStr)
		weights = append(weights, weight)
	}
	sort.Ints(weights)
	bs, _, _ = rd.ReadLine()
	m, _ := strconv.Atoi(string(bs))
	ans := 0
	var dfs func(nowIndex, nowSum int)
	dfs = func(nowIndex, nowSum int) {
		if nowSum > m {
			return
		}
		if nowSum+weights[0]-1 >= m {
			ans++
			return
		}
		for i := nowIndex; i < len(weights); i++ {
			dfs(i, nowSum+weights[i])
		}
	}
	dfs(0, 0)
	fmt.Println(ans)
}
