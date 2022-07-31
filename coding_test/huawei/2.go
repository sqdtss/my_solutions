package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	var M, N int
	fmt.Scan(&M, &N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&P[i])
	}
	sort.Ints(P)
	type data struct {
		Q, S int
	}
	datas := make([]data, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&datas[i].Q)
	}
	for i := 0; i < M; i++ {
		fmt.Scan(&datas[i].S)
	}
	sort.Slice(datas, func(i, j int) bool {
		if datas[i].S != datas[j].S {
			return datas[i].S > datas[j].S
		}
		return datas[i].Q > datas[j].Q
	})
	anses := make([]int, N)
	for j := 0; j < M; j++ {
		nowMin := math.MaxInt64
		nowIdx := 0
		for i := 0; i < N; i++ {
			if datas[j].Q <= P[i] && anses[i]+datas[j].S < nowMin {
				nowMin = anses[i] + datas[j].S
				nowIdx = i
			}
		}
		anses[nowIdx] += datas[j].S
	}
	fmt.Println(max(anses...))
}

func max(nums ...int) int {
	maxE := nums[0]
	for _, num := range nums {
		if num > maxE {
			maxE = num
		}
	}
	return maxE
}
