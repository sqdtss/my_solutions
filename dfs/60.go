package solution

import "strconv"

// 60. 排列序列
// 给出集合 [1,2,3,...,n]，其所有元素共有 n! 种排列。
// 按大小顺序列出所有排列情况，并一一标记，当 n = 3 时, 所有排列如下：
// "123"
// "132"
// "213"
// "231"
// "312"
// "321"
// 给定 n 和 k，返回第 k 个排列。

// 示例 1：
// 输入：n = 3, k = 3
// 输出："213"

// 示例 2：
// 输入：n = 4, k = 9
// 输出："2314"

// 示例 3：
// 输入：n = 3, k = 1
// 输出："123"

// 提示：
// 1 <= n <= 9
// 1 <= k <= n!

func getPermutation(n int, k int) (res string) {
	factorial := make([]int, n+1) // 阶乘数组
	factorial[0] = 1
	for i := 1; i <= n; i++ {
		factorial[i] = factorial[i-1] * i
	}
	visited := make([]bool, n+1) // dfs所用visit数组
	var dfs func(int)
	dfs = func(index int) {
		if index == n {
			return
		}
		// 计算还未确定的数字的全排列的个数，第 1 次进入的时候是 n - 1
		cnt := factorial[n-1-index]
		for i := 1; i <= n; i++ {
			if visited[i] {
				continue
			}
			if cnt < k {
				k -= cnt
				continue
			}
			res += strconv.Itoa(i)
			visited[i] = true
			dfs(index + 1)
			// 注意 1：不可以回溯（重置变量），算法设计是「一下子来到叶子结点」，没有回头的过程
			// 注意 2：这里要加 return，后面的数没有必要遍历去尝试了
			return
		}
	}
	dfs(0)
	return
}
