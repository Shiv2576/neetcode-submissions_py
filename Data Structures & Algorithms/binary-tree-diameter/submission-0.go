/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


var diameter int

func diameterOfBinaryTree(root *TreeNode) int {
    diameter = 0
    height(root)
    return diameter
}


func height(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := height(root.Left)
    right := height(root.Right)

    diameter = max(diameter , left + right)


    return 1 + max(right, left)
}

func max(a, b int) int {
    if a >= b {
        return a
    }

    return b
}