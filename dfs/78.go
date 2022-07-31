package solution

// 78. 子集
// 给你一个整数数组 nums ，数组中的元素 互不相同 。返回该数组所有可能的子集（幂集）。
// 解集 不能 包含重复的子集。你可以按 任意顺序 返回解集。

// 示例 1：
// 输入：nums = [1,2,3]
// 输出：[[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]

// 示例 2：
// 输入：nums = [0]
// 输出：[[],[0]]

// 提示：
// 1 <= nums.length <= 10
// -10 <= nums[i] <= 10
// nums 中的所有元素 互不相同

var res [][]int
var visited []bool

func subsets(nums []int) [][]int {
	res = nil
	visited = make([]bool, len(nums)+1)
	dfs(nil, nums, 0, 0, len(nums))
	return res
}

func dfs(nowSz, nums []int, index, step, fullStep int) {
	if step == fullStep {
		tmp := make([]int, len(nowSz))
		copy(tmp, nowSz)
		res = append(res, tmp)
		return
	}
	for i := index; i <= len(nums); i++ {
		if visited[i] {
			continue
		}
		if i == len(nums) { // 用i == len(nums)来代表啥都不加的情况，这种情况visited不修改，且index不修改
			dfs(nowSz, nums, i, step+1, fullStep)
		} else {
			visited[i] = true
			tmp := append(nowSz, nums[i])
			dfs(tmp, nums, i+1, step+1, fullStep)
			visited[i] = false
		}
	}
}

// my method 2：
//func subsets(nums []int) (ans [][]int) {
//	ans = append(ans, nil)
//	var dfs func(int, []int)
//	dfs = func(nowIndex int, nowNums []int) {
//		if nowIndex == len(nums) {
//			tmp := make([]int, len(nowNums))
//			copy(tmp, nowNums)
//			ans = append(ans, tmp)
//			return
//		}
//		nowNums = append(nowNums, nums[nowIndex])
//		for i := nowIndex + 1; i <= len(nums); i++ {
//			dfs(i, nowNums)
//		}
//	}
//	for i := 0; i < len(nums); i++ {
//		dfs(i, nil)
//	}
//	return
//}
