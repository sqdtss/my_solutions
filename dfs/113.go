package solution

//113. 路径总和 II
//给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
//叶子节点 是指没有子节点的节点。
//
//示例 1：
//输入：root = [5,4,8,11,null,13,4,7,2,null,null,5,1], targetSum = 22
//输出：[[5,4,11,2],[5,8,4,5]]
//
//示例 2：
//输入：root = [1,2,3], targetSum = 5
//输出：[]
//
//示例 3：
//输入：root = [1,2], targetSum = 0
//输出：[]
//
//提示：
//树中节点总数在范围 [0, 5000] 内
//-1000 <= Node.val <= 1000
//-1000 <= targetSum <= 1000s

func pathSum(root *TreeNode, targetSum int) (ans [][]int) {
	var dfs func(*TreeNode, []int, int)
	dfs = func(node *TreeNode, nowSumSz []int, nowSum int) {
		if node == nil {
			return
		}
		nowSumSz = append(nowSumSz, node.Val)
		nowSum += node.Val
		if node.Left == nil && node.Right == nil {
			if nowSum == targetSum {
				tmp := make([]int, len(nowSumSz))
				copy(tmp, nowSumSz)
				ans = append(ans, tmp)
			}
			return
		}
		dfs(node.Left, nowSumSz, nowSum)
		dfs(node.Right, nowSumSz, nowSum)
	}
	dfs(root, nil, 0)
	return
}
