package solution

import (
	"sort"
)

// 435. 无重叠区间
// 给定一个区间的集合，找到需要移除区间的最小数量，使剩余区间互不重叠。
// 注意:
// 可以认为区间的终点总是大于它的起点。
// 区间 [1,2] 和 [2,3] 的边界相互“接触”，但没有相互重叠。

// 示例 1:
// 输入: [ [1,2], [2,3], [3,4], [1,3] ]
// 输出: 1
// 解释: 移除 [1,3] 后，剩下的区间没有重叠。

// 示例 2:
// 输入: [ [1,2], [1,2], [1,2] ]
// 输出: 2
// 解释: 你需要移除两个 [1,2] 来使剩下的区间没有重叠。

// 示例 3:
// 输入: [ [1,2], [2,3] ]
// 输出: 0
// 解释: 你不需要移除任何区间，因为它们已经是无重叠的了。

// 方法一: dp 此方法会超时
// func eraseOverlapIntervals(intervals [][]int) int {
// 	sort.Slice(intervals, func(i, j int) bool {
// 		return intervals[i][0] < intervals[j][0]
// 	})
// 	dp := make([]int, len(intervals))
// 	for i := range dp {
// 		dp[i] = 1
// 	}
// 	for i := 0; i < len(intervals); i++ {
// 		for j := 0; j < i; j++ {
// 			if intervals[i][0] >= intervals[j][1] {
// 				dp[i] = max(dp[i], dp[j]+1)
// 			}
// 		}
// 	}
// 	return len(intervals) - max(dp...)
// }

// func max(a ...int) int {
// 	res := a[0]
// 	for _, v := range a[1:] {
// 		if v > res {
// 			res = v
// 		}
// 	}
// 	return res
// }

// 方法二: 用右边界排序
func eraseOverlapIntervals(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	ans := 1
	right := intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] >= right {
			ans++
			right = intervals[i][1]
		}
	}
	return len(intervals) - ans
}
