/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func maxDepth(root *TreeNode) int {
    // time complexity : O(n)
    // if root == nil {
    //     return 0
    // }
    // return 1+ max(maxDepth(root.Left), maxDepth(root.Right))
    if root == nil {
        return 0
    }
    
    // bfs
    queue := []*TreeNode{root}
    var depth int
    for len(queue) > 0 {
        depth++
        size := len(queue)
        for i:=0; i<size; i++ {
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