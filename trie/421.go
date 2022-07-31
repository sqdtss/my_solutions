package solution

//421. 数组中两个数的最大异或值
//给你一个整数数组 nums ，返回 nums[i] XOR nums[j] 的最大运算结果，其中 0 ≤ i ≤ j < n 。
//进阶：你可以在 O(n) 的时间解决这个问题吗？
//
//示例 1：
//输入：nums = [3,10,5,25,2,8]
//输出：28
//解释：最大运算结果是 5 XOR 25 = 28.
//
//示例 2：
//输入：nums = [0]
//输出：0
//
//示例 3：
//输入：nums = [2,4]
//输出：6
//
//示例 4：
//输入：nums = [8,10,2]
//输出：10
//
//示例 5：
//输入：nums = [14,70,53,83,49,91,36,80,92,51,66,70]
//输出：127
//
//提示：
//1 <= nums.length <= 2 * 104
//0 <= nums[i] <= 231 - 1
//
//const highBit = 30
//
//type trie struct {
//	left, right *trie
//}
//
//func (t *trie) add(num int) {
//	cur := t
//	for i := highBit; i >= 0; i-- {
//		bit := num >> i & 1
//		if bit == 0 {
//			if cur.left == nil {
//				cur.left = &trie{}
//			}
//			cur = cur.left
//		} else {
//			if cur.right == nil {
//				cur.right = &trie{}
//			}
//			cur = cur.right
//		}
//	}
//}
//
//func (t *trie) check(num int) (x int) {
//	cur := t
//	for i := highBit; i >= 0; i-- {
//		bit := num >> i & 1
//		if bit == 0 {
//			// a_i 的第 k 个二进制位为 0，应当往表示 1 的子节点 right 走
//			if cur.right != nil {
//				cur = cur.right
//				x = x*2 + 1
//			} else {
//				cur = cur.left
//				x = x * 2
//			}
//		} else {
//			// a_i 的第 k 个二进制位为 1，应当往表示 0 的子节点 left 走
//			if cur.left != nil {
//				cur = cur.left
//				x = x*2 + 1
//			} else {
//				cur = cur.right
//				x = x * 2
//			}
//		}
//	}
//	return
//}
//
//func findMaximumXOR(nums []int) (x int) {
//	root := &trie{}
//	for i := 1; i < len(nums); i++ {
//		// 将 nums[i-1] 放入字典树，此时 nums[0 .. i-1] 都在字典树中
//		root.add(nums[i-1])
//		// 将 nums[i] 看作 ai，找出最大的 x 更新答案
//		x = max(x, root.check(nums[i]))
//	}
//	return
//}
//
//func max(param ...int) int {
//	maxE := param[0]
//	for i := 1; i < len(param); i++ {
//		if maxE < param[i] {
//			maxE = param[i]
//		}
//	}
//	return maxE
//}
