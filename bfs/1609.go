package bfs

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	var queue []*TreeNode
	queue = append(queue, root)
	isEven := true
	for len(queue) != 0 {
		if isEven {
			for i := range queue {
				if queue[i].Val%2 != 1 {
					return false
				}
				if i > 0 && (queue[i].Val <= queue[i-1].Val) {
					return false
				}
			}
		} else {
			for i := range queue {
				if queue[i].Val%2 != 0 {
					return false
				}
				if i > 0 && (queue[i].Val >= queue[i-1].Val) {
					return false
				}
			}
		}
		var queueTemp []*TreeNode
		for _, q := range queue {
			if q.Left != nil {
				queueTemp = append(queueTemp, q.Left)
			}
			if q.Right != nil {
				queueTemp = append(queueTemp, q.Right)
			}
		}
		queue = queueTemp
		isEven = !isEven
	}
	return true
}
