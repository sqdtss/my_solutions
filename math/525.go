package solution

// 前缀和 + hash
// 由于「0 和 1 的数量相同」等价于「1 的数量减去 0 的数量等于 0」，我们可以将数组中的 0 视作 −1，则原问题转换成「求最长的连续子数组，其元素和为 0」。

// 525. 连续数组
// 给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
 
// 示例 1:
// 输入: nums = [0,1]
// 输出: 2
// 说明: [0, 1] 是具有相同数量 0 和 1 的最长连续子数组。

// 示例 2:
// 输入: nums = [0,1,0]
// 输出: 2
// 说明: [0, 1] (或 [1, 0]) 是具有相同数量0和1的最长连续子数组。
 
// 提示：
// 1 <= nums.length <= 105
// nums[i] 不是 0 就是 1

func findMaxLength(nums []int) (ans int) {
	counter := 0
	m := map[int]int{0: -1}
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if preIndex, ok := m[counter]; ok {
			ans = Max(ans, i-preIndex)
		} else {
			m[counter] = i
		}
	}
	return
}

func Max(param ...int) int {
	maxE := param[0]
	for i := 1; i < len(param); i++ {
		if maxE < param[i] {
			maxE = param[i]
		}
	}
	return maxE
}