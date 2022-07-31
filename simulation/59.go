package simulation

//59. 螺旋矩阵 II
//给你一个正整数 n ，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的 n x n 正方形矩阵 matrix 。
//
//示例 1：
//输入：n = 3
//输出：[[1,2,3],[8,9,4],[7,6,5]]
//
//示例 2：
//输入：n = 1
//输出：[[1]]
//
//提示：
//1 <= n <= 20

func generateMatrix(n int) (ans [][]int) {
	sz := make([][]int, n)
	for i := range sz {
		sz[i] = make([]int, n)
	}
	count := 1
	left, right, top, down := 0, n-1, 0, n-1
	for left <= right && top <= down {
		for i := left; i <= right; i++ {
			sz[top][i] = count
			count++
		}
		top++
		if top > down {
			break
		}
		for i := top; i <= down; i++ {
			sz[i][right] = count
			count++
		}
		right--
		if left > right {
			break
		}
		for i := right; i >= left; i-- {
			sz[down][i] = count
			count++
		}
		down--
		if top > down {
			break
		}
		for i := down; i >= top; i-- {
			sz[i][left] = count
			count++
		}
		left++
		if left > right {
			break
		}
	}
	for _, tmp := range sz {
		ans = append(ans, tmp)
	}
	return
}
