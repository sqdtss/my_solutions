package solution

import "container/heap"

// 373. 查找和最小的K对数字
// 给定两个以升序排列的整数数组 nums1 和 nums2 , 以及一个整数 k 。
// 定义一对值 (u,v)，其中第一个元素来自 nums1，第二个元素来自 nums2 。
// 请找到和最小的 k 个数对 (u1,v1),  (u2,v2)  ...  (uk,vk) 。

// 示例 1:
// 输入: nums1 = [1,7,11], nums2 = [2,4,6], k = 3
// 输出: [1,2],[1,4],[1,6]
// 解释: 返回序列中的前 3 对数：
//      [1,2],[1,4],[1,6],[7,2],[7,4],[11,2],[7,6],[11,4],[11,6]

// 	 示例 2:
// 输入: nums1 = [1,1,2], nums2 = [1,2,3], k = 2
// 输出: [1,1],[1,1]
// 解释: 返回序列中的前 2 对数：
//      [1,1],[1,1],[1,2],[2,1],[1,2],[2,2],[1,3],[1,3],[2,3]

// 示例 3:
// 输入: nums1 = [1,2], nums2 = [3], k = 3
// 输出: [1,3],[2,3]
// 解释: 也可能序列中所有的数对都被返回:[1,3],[2,3]

// 提示:
// 1 <= nums1.length, nums2.length <= 104
// -109 <= nums1[i], nums2[i] <= 109
// nums1, nums2 均为升序排列
// 1 <= k <= 1000

func kSmallestPairs(nums1 []int, nums2 []int, k int) (res [][]int) {
	h := hp{}
	for i := 0; i < len(nums1) && i < k; i++ {
		for j := 0; j < len(nums2) && j < k; j++ {
			heap.Push(&h, pair{nums1[i], nums2[j]})
		}
	}
	for i := 0; i < k && len(h) > 0; i++ {
		tmp := heap.Pop(&h).(pair)
		res = append(res, []int{tmp.num1, tmp.num2})
	}
	return
}

type pair struct {
	num1, num2 int
}
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].num1+h[i].num2 < h[j].num1+h[j].num2
}

func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
