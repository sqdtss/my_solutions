package solution

//653. 两数之和 IV - 输入 BST
//给定一个二叉搜索树 root 和一个目标结果 k，如果 BST 中存在两个元素且它们的和等于给定的目标结果，则返回 true。
//
//示例 1：
//输入: root = [5,3,6,2,4,null,7], k = 9
//输出: true
//
//示例 2：
//输入: root = [5,3,6,2,4,null,7], k = 28
//输出: false
//
//提示:
//二叉树的节点个数的范围是  [1, 104].
//-104 <= Node.val <= 104
//root 为二叉搜索树
//-105 <= k <= 105

//type TreeNode struct {
//	Val   int
//	Left  *TreeNode
//	Right *TreeNode
//}
//
// 先序遍历 + hashmap
//func findTarget(root *TreeNode, k int) bool {
//	mp := map[int]bool{}
//	var dfs func(node *TreeNode) bool
//	dfs = func(node *TreeNode) bool {
//		if node == nil {
//			return false
//		}
//		if mp[k-node.Val] {
//			return true
//		}
//		mp[node.Val] = true
//		return dfs(node.Left) || dfs(node.Right)
//	}
//	return dfs(root)
//}
