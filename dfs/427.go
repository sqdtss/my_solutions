package solution

//427. 建立四叉树
//https://leetcode-cn.com/problems/construct-quad-tree/

//func construct(grid [][]int) *Node {
//	var dfs func(int, int, int, int) *Node
//	dfs = func(top, down, left, right int) *Node {
//		for i := top; i < down; i++ {
//			for j := left; j < right; j++ {
//				if grid[i][j] != grid[top][left] {
//					return &Node{
//						true,
//						false,
//						dfs(top, (top+down)/2, left, (left+right)/2),
//						dfs(top, (top+down)/2, (left+right)/2, right),
//						dfs((top+down)/2, down, left, (left+right)/2),
//						dfs((top+down)/2, down, (left+right)/2, right),
//					}
//				}
//			}
//		}
//		return &Node{Val: grid[top][left] == 1, IsLeaf: true}
//	}
//	return dfs(0, len(grid), 0, len(grid))
//}
