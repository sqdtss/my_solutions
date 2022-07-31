package solution

func combinationSum(candidates []int, target int) (ans [][]int) {
	var dfs func(nowSum, nowIdx int, nowList []int)
	dfs = func(nowSum, nowIdx int, nowList []int) {
		if nowSum == target {
			ans = append(ans, nowList)
			return
		}

		if nowSum > target {
			return
		}

		for idx := nowIdx; idx < len(candidates); idx++ {
			tmpList := append([]int{}, nowList...)
			tmpList = append(tmpList, candidates[idx])
			dfs(nowSum+candidates[idx], idx, tmpList)
		}
	}

	dfs(0, 0, []int{})
	return
}
