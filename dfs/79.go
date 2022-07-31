package solution

// 79. 单词搜索
// 给定一个 m x n 二维字符网格 board 和一个字符串单词 word 。如果 word 存在于网格中，返回 true ；否则，返回 false 。
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。

// 示例 1：
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
// 输出：true

// 示例 2：
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "SEE"
// 输出：true
// 示例 3：
// 输入：board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCB"
// 输出：false

// 提示：
// m == board.length
// n = board[i].length
// 1 <= m, n <= 6
// 1 <= word.length <= 15
// board 和 word 仅由大小写英文字母组成

// 测试
// import "fmt"

// func main() {
// 	fmt.Println(exist([][]byte{[]byte("a")}, "a"))
// 	fmt.Println(exist([][]byte{[]byte("abce"), []byte("sfcs"), []byte("adee")}, "abcb"))
// 	fmt.Println(exist([][]byte{[]byte("abce"), []byte("sfcs"), []byte("adee")}, "abccfs"))
// }

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	direction := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	var dfs func(int, int, int) bool
	dfs = func(i, j, index int) bool {
		if board[i][j] == word[index] {
			if index+1 == len(word) {
				return true
			}
			for _, direct := range direction {
				newI, newJ := i+direct[0], j+direct[1]
				if newI < 0 || newJ < 0 || newI >= m || newJ >= n || visited[newI][newJ] {
					continue
				}
				visited[newI][newJ] = true
				if dfs(newI, newJ, index+1) {
					return true
				}
				visited[newI][newJ] = false
			}
		}
		return false
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			visited[i][j] = true
			if dfs(i, j, 0) {
				return true
			}
			visited[i][j] = false
		}
	}
	return false
}
