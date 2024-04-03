/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func countNodes(root *TreeNode) int {
    return dfs(root, 0)
}

func dfs(root *TreeNode, currCnt int) int {
    if root == nil {
        return 0
    }

    if root.Left != nil {
        currCnt += 1
    }

    if root.Right != nil {
        currCnt += 1
    }

    return 1 + dfs(root.Left, currCnt) + dfs(root.Right, currCnt)
} 