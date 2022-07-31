package bfs

//542. 01 矩阵
//给定一个由 0 和 1 组成的矩阵 matrix ，请输出一个大小相同的矩阵，其中每一个格子是 matrix 中对应位置元素到最近的 0 的距离。
//两个相邻元素间的距离为 1 。
//
//示例 1：
//输入：matrix = [[0,0,0],[0,1,0],[0,0,0]]
//输出：[[0,0,0],[0,1,0],[0,0,0]]
//
//示例 2：
//输入：matrix = [[0,0,0],[0,1,0],[1,1,1]]
//输出：[[0,0,0],[0,1,0],[1,2,1]]
//
//提示：
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 104
//1 <= m * n <= 104
//matrix[i][j] is either 0 or 1.
//matrix 中至少有一个 0

func updateMatrix(matrix [][]int) [][]int {
	m, n := len(matrix), len(matrix[0])
	direction := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	ans := make([][]int, m)
	visited := make([][]bool, m)
	for i := range ans {
		ans[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}
	var queue [][]int
	for i := range matrix {
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				queue = append(queue, []int{i, j})
				visited[i][j] = true
			}
		}
	}
	for len(queue) > 0 {
		front := queue[0]
		queue = queue[1:]
		for _, direct := range direction {
			newX, newY := front[0]+direct[0], front[1]+direct[1]
			if newX < 0 || newX >= m || newY < 0 || newY >= n || visited[newX][newY] {
				continue
			}
			queue = append(queue, []int{newX, newY})
			ans[newX][newY] = ans[front[0]][front[1]] + 1
			visited[newX][newY] = true
		}
	}
	return ans
}
