/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    // recursive dfs
    /*if root == nil {
        return 0
    }

    return 1+max(maxDepth(root.Left), maxDepth(root.Right))*/

    // BFS
    if root == nil {
        return 0
    }

    queue := []*TreeNode{root}
    depth := 0

    for len(queue) > 0 {
        depth++

        level := len(queue)

        for i := 0; i < level; i++ {
            curr := queue[0]
            queue = queue[1:]

            if curr.Left != nil {
                queue = append(queue, curr.Left)
            }

            if curr.Right != nil {
                queue = append(queue, curr.Right)
            }
        }
    }

    return depth
}

