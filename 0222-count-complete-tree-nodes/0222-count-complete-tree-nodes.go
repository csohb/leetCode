/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func countNodes(root *TreeNode) int {
    left := leftHeight(root)
    right := rightHeight(root)

    if left == right {
        return int(math.Pow(2, float64(left))) - 1
    } else {
        return 1 + countNodes(root.Left) + countNodes(root.Right)
    }
}


func leftHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + leftHeight(root.Left)
}

func rightHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + rightHeight(root.Right)
}