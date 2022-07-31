package solution

// 全排列

func permute(nums []int) (ans [][]int) {
	visited := make([]bool, len(nums))
	var dfs func([]int)
	dfs = func(nowSz []int) {
		if len(nowSz) == len(nums) {
			tmp := make([]int, len(nowSz))
			copy(tmp, nowSz)
			ans = append(ans, tmp)
			return
		}
		for i, num := range nums {
			if !visited[i] {
				tmp := append(nowSz, num)
				visited[i] = true
				dfs(tmp)
				visited[i] = false
			}
		}
	}
	dfs(nil)
	return
}
