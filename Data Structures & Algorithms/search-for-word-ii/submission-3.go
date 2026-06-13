type TreeNode struct {
    children [26]*TreeNode
    word string
}

func findWords(board [][]byte, words []string) []string {
    root := &TreeNode{}

    for _, word := range words {
        insert(root, word)
    }

    res := []string{}
    m, n := len(board), len(board[0])

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            dfs(board, i, j, root, &res)
        }
    }

    return res
}


func insert(root *TreeNode , word string) {
    for _ ,ch := range word {
        idx := ch - 'a'
        if root.children[idx] == nil {
            root.children[idx] = &TreeNode{}
        }
        root = root.children[idx]
    }

    root.word = word
}


func dfs(board [][]byte, i, j int, node *TreeNode, res *[]string) {
    ch := board[i][j]

    if ch == '#' {
        return
    }

    idx := ch - 'a'
    if node.children[idx] == nil {
        return
    }

    node = node.children[idx]

    if node.word != "" {
        *res = append(*res, node.word)
        node.word = "" // avoid duplicates
    }

    board[i][j] = '#'

    dirs := [][]int{{1,0},{-1,0},{0,1},{0,-1}}
    for _, d := range dirs {
        ni, nj := i+d[0], j+d[1]
        if ni >= 0 && ni < len(board) && nj >= 0 && nj < len(board[0]) {
            dfs(board, ni, nj, node, res)
        }
    }

    board[i][j] = ch
}
