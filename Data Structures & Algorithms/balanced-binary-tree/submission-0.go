/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isBalanced(root *TreeNode) bool {
    return height(root) != -1
}


func height(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := height(root.Left)
    if left == -1 {
        return -1
    }


    right := height(root.Right)
    if right == -1 {
        return -1
    }


    if abs(right - left) > 1 {
        return -1
    }


    return max(left , right) + 1
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}