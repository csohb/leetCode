/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var target int 
func hasPathSum(root *TreeNode, targetSum int) bool {
    // is there's path that maked targetSum, then return true
    target = targetSum
    return dfs(root, 0) 
}

func dfs(root *TreeNode, currSum int) bool {
    if root == nil {
        return false
    }

    if root.Left == nil && root.Right == nil {
        return (root.Val + currSum) == target
    }

    currSum += root.Val
    left := dfs(root.Left, currSum)
    right := dfs(root.Right, currSum)

    return left || right

}

