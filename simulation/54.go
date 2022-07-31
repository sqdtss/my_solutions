package simulation

//54. 螺旋矩阵
//给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
//
//示例 1：
//输入：matrix = [[1,2,3],[4,5,6],[7,8,9]]
//输出：[1,2,3,6,9,8,7,4,5]
//
//示例 2：
//输入：matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]
//输出：[1,2,3,4,8,12,11,10,9,5,6,7]
//
//提示：
//m == matrix.length
//n == matrix[i].length
//1 <= m, n <= 10
//-100 <= matrix[i][j] <= 100

func spiralOrder(matrix [][]int) (ans []int) {
	row, column := len(matrix), len(matrix[0])
	left, right, up, down := 0, column-1, 0, row-1
	for {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[up][i])
		}
		up++
		if up >= down {
			break
		}
		for i := up; i <= down; i++ {
			ans = append(ans, matrix[i][right])
		}
		right--
		if left >= right {
			break
		}
		for i := right; i >= left; i-- {
			ans = append(ans, matrix[down][i])
		}
		down--
		if up >= down {
			break
		}
		for i := down; i >= up; i-- {
			ans = append(ans, matrix[i][left])
		}
		left++
		if left >= right {
			break
		}
	}
	return
}
