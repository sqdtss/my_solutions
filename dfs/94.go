package solution

//94. 二叉树的中序遍历
//给定一个二叉树的根节点 root ，返回 它的 中序 遍历 。
//示例 1：
//输入：root = [1,null,2,3]
//输出：[1,3,2]
//
//示例 2：
//输入：root = []
//输出：[]
//
//示例 3：
//输入：root = [1]
//输出：[1]
//
//提示：
//树中节点数目在范围 [0, 100] 内
//-100 <= Node.val <= 100
//
//进阶: 递归算法很简单，你可以通过迭代算法完成吗？

// 递归
func inorderTraversal(root *TreeNode) (ans []int) {
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		ans = append(ans, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return
}

// 迭代
//func inorderTraversal(root *TreeNode) (ans []int) {
//	if root == nil {
//		return
//	}
//	var s []*TreeNode
//	node := root
//	for node != nil || len(s) > 0 {
//		for node != nil {
//			s = append(s, node)
//			node = node.Left
//		}
//		node = s[len(s)-1]
//		s = s[:len(s)-1]
//		ans = append(ans, node.Val)
//		node = node.Right
//	}
//	return
//}
