package solution

//230. 二叉搜索树中第K小的元素
//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
//示例 1：
//输入：root = [3,1,4,null,2], k = 1
//输出：1
//
//示例 2：
//输入：root = [5,3,6,2,4,null,null,1], k = 3
//输出：3
//
//提示：
//树中的节点数为 n 。
//1 <= k <= n <= 104
//0 <= Node.val <= 104
//
//进阶：如果二叉搜索树经常被修改（插入/删除操作）并且你需要频繁地查找第 k 小的值，你将如何优化算法

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}

// 方法一：中序遍历
//func kthSmallest(root *TreeNode, k int) (ans int) {
//	var dfs func(*TreeNode)
//	dfs = func(node *TreeNode) {
//		if node == nil || k == 0 {
//			return
//		}
//		dfs(node.Left)
//		k--
//		if k == 0 {
//			ans = node.Val
//		}
//		dfs(node.Right)
//	}
//	dfs(root)
//	return
//}

// 方法二
//func kthSmallest(root *TreeNode, k int) int {
//	ans := -1
//	var dfs func(*TreeNode, int) int
//	dfs = func(node *TreeNode, k int) int {
//		if node == nil || ans != -1 {
//			return 0
//		}
//		left := dfs(node.Left, k)
//		if left+1 == k {
//			ans = node.Val
//		}
//		right := dfs(node.Right, k-left-1)
//		return left + right + 1
//	}
//	dfs(root, k)
//	return ans
//}
