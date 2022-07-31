package solution

import "sort"

// 全排列 + 去重

func permuteUnique(nums []int) (ans [][]int) {
	sort.Ints(nums)
	mp := map[int]bool{}
	visited := make([]bool, len(nums))
	var dfs func(nowSz []int, nowLen int)
	dfs = func(nowSz []int, nowLen int) {
		if nowLen == len(nums) {
			tmp := make([]int, nowLen)
			copy(tmp, nowSz)
			ans = append(ans, tmp)
			return
		}
		for i, num := range nums {
			if i > 0 && num == nums[i-1] && !mp[i-1] {
				continue
			}
			if !visited[i] {
				visited[i] = true
				mp[i] = true
				tmp := append(nowSz, num)
				dfs(tmp, nowLen+1)
				mp[i] = false
				visited[i] = false
			}
		}
	}
	dfs(nil, 0)
	return
}
