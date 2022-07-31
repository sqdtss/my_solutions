package solution

// dfs + trie

// 212. 单词搜索 II
// 给定一个 m x n 二维字符网格 board 和一个单词（字符串）列表 words，找出所有同时在二维网格和字典中出现的单词。
// 单词必须按照字母顺序，通过 相邻的单元格 内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母在一个单词中不允许被重复使用。

// 示例 1：
// 输入：board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
// 输出：["eat","oath"]

// 示例 2：
// 输入：board = [["a","b"],["c","d"]], words = ["abcb"]
// 输出：[]

// 提示：
// m == board.length
// n == board[i].length
// 1 <= m, n <= 12
// board[i][j] 是一个小写英文字母
// 1 <= words.length <= 3 * 104
// 1 <= words[i].length <= 10
// words[i] 由小写英文字母组成
// words 中的所有字符串互不相同

type Trie struct {
	children [26]*Trie
	word     string
}

func (this *Trie) Insert(word string) {
	for _, c := range word {
		index := c - 'a'
		if this.children[index] == nil {
			this.children[index] = &Trie{}
		}
		this = this.children[index]
	}
	this.word = word
}

func findWords(board [][]byte, words []string) []string {
	directions := [][]int{{0, 1}, {1, 0}, {-1, 0}, {0, -1}}
	t := &Trie{}
	for _, word := range words {
		t.Insert(word)
	}
	m, n := len(board), len(board[0])
	seen := map[string]bool{}
	var dfs func(node *Trie, x, y int)
	dfs = func(node *Trie, x, y int) {
		c := board[x][y]
		node = node.children[c-'a']
		if node == nil {
			return
		}

		if node.word != "" {
			seen[node.word] = true
		}

		board[x][y] = '#'
		for _, direction := range directions {
			nx, ny := x+direction[0], y+direction[1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && board[nx][ny] != '#' {
				dfs(node, nx, ny)
			}
		}
		board[x][y] = c
	}
	for i, row := range board {
		for j := range row {
			dfs(t, i, j)
		}
	}

	ans := make([]string, 0, len(seen))
	for s := range seen {
		ans = append(ans, s)
	}
	return ans
}
