/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
import (
    "fmt"
)

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    } else if root.Left == nil && root.Right == nil {
        return 1
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func dfs(root *TreeNode, result int) int {
    if root == nil {
        return result
    }

    if root.Left == nil && root.Right == nil {
        return result
    }

    if root.Left != nil {
        result++
        dfs(root.Left, result)
    }

    if root.Right != nil {
        result++
        dfs(root.Right, result)
    }
    return result
}

func dfsLeft(root *TreeNode, result int) int {
    if root.Left == nil {
        return result
    }

    if root.Left != nil {
        result++
        dfs(root.Left, result)
    }
    return result
}

func dfsRight(root *TreeNode, result int) int {
    if root.Right == nil {
        return result
    }

    if root.Right != nil {
        result++
        dfs(root.Right, result)
    }
    return result
}