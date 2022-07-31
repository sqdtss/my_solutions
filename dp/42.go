package solution

//42. 接雨水
//给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。
//
//示例 1：
//输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
//输出：6
//解释：上面是由数组 [0,1,0,2,1,0,1,3,2,1,2,1] 表示的高度图，在这种情况下，可以接 6 个单位的雨水（蓝色部分表示雨水）。
//
//示例 2：
//输入：height = [4,2,0,3,2,5]
//输出：9
//
//提示：
//n == height.length
//1 <= n <= 2 * 104
//0 <= height[i] <= 105

func trap(height []int) (ans int) {
	lenHeight := len(height)
	maxSzLeft, maxSzRight := make([]int, lenHeight), make([]int, lenHeight)
	maxSzLeft[0] = height[0]
	maxSzRight[lenHeight-1] = height[lenHeight-1]
	for i := 1; i < lenHeight; i++ {
		maxSzLeft[i] = max(height[i], maxSzLeft[i-1])
	}
	for i := lenHeight - 2; i >= 0; i-- {
		maxSzRight[i] = max(height[i], maxSzRight[i+1])
	}
	for i := 0; i < lenHeight; i++ {
		ans += max(min(maxSzLeft[i], maxSzRight[i])-height[i], 0)
	}
	return
}

//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
//
//func min(a, b int) int {
//	if a < b {
//		return a
//	}
//	return b
//}
