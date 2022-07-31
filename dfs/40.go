package solution

// dfs + 去重
import (
	"sort"
)

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	mp := map[int]bool{}
	var dfs func(int, int, []int)
	dfs = func(startIndex, nowSum int, nowSz []int) {
		if nowSum == target {
			tmp := make([]int, len(nowSz))
			copy(tmp, nowSz)
			ans = append(ans, tmp)
			return
		} else if nowSum > target {
			return
		}
		for i := startIndex; i < len(candidates); i++ {
			if i > 0 && candidates[i] == candidates[i-1] && !mp[i-1] {
				continue
			}
			tmp := append(nowSz, candidates[i])
			mp[i] = true
			dfs(i+1, nowSum+candidates[i], tmp)
			mp[i] = false
		}
	}
	dfs(0, 0, nil)
	return
}
