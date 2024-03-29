/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    // left <-> right
    if root == nil {
        return nil
    }

    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    for len(queue) > 0 {
        start := queue[0]
        queue = queue[1:]

        start.Left, start.Right = start.Right, start.Left
        if start.Left != nil {
            queue = append(queue, start.Left)
        } 

        if start.Right != nil {
            queue = append(queue, start.Right)
        }
    }

    return root
}

