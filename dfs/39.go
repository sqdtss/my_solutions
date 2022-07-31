package solution

func combinationSum(candidates []int, target int) (res [][]int) {
	var dfs func(startIndex int, nowSz []int, nowSum int, candidates []int, target int)
	dfs = func(startIndex int, nowSz []int, nowSum int, candidates []int, target int) {
		if nowSum == target {
			tmp := make([]int, len(nowSz))
			copy(tmp, nowSz)
			res = append(res, tmp)
			return
		} else if nowSum > target {
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			tmp := append(nowSz, candidates[i])
			dfs(i, tmp, nowSum+candidates[i], candidates, target)
		}
	}
	dfs(0, nil, 0, candidates, target)
	return res
}
