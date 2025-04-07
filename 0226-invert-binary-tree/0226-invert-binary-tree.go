/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    // BFS
    // switch in each level of tree
    if root == nil {
        return nil
    }

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        level := len(queue)

        for i:=0; i < level; i++ {
            curr := queue[0]
            queue = queue[1:]

            curr.Left, curr.Right = curr.Right, curr.Left

            if curr.Left != nil {
                queue = append(queue, curr.Left)
            }

            if curr.Right != nil {
                queue = append(queue, curr.Right)
            }
        }
    }

    return root
}