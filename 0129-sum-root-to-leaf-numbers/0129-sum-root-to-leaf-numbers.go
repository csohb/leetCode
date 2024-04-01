/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
    return getSum(root,0)
}

func getSum(node *TreeNode, num int) int {
    if node == nil {
        return 0
    }

    num = num * 10 + node.Val
    if node.Left == nil && node.Right == nil {
        return num
    }

    return getSum(node.Left, num) + getSum(node.Right, num)
}