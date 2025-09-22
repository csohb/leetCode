/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return nil
    }
    res := [][]int{}

    queue := []*TreeNode{root}

    for len(queue) > 0 {
        size := len(queue)
        level := []int{}
        for i:=0; i<size; i++ {
            n:= queue[0]
            queue = queue[1:]

            if n.Left != nil {
                queue = append(queue, n.Left)
            }

            if n.Right != nil {
                queue = append(queue, n.Right)
            }
            level = append(level, n.Val)
        }
        res = append(res, level)
    }

    return res
}