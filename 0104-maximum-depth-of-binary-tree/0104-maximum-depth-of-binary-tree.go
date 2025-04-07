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
    /*if root == nil {
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

    return depth*/

    // DFS

    if root == nil {
        return 0
    }
    stack := []Depth{{root,1}}
    maxDepth := 0 

    for len(stack) > 0 {
        curr := stack[len(stack)-1]
        stack = stack[:len(stack)-1]

        if maxDepth < curr.depth {
            maxDepth = curr.depth
        }

        if curr.node.Right != nil {
            stack = append(stack, Depth{curr.node.Right,curr.depth+1})
        }

        if curr.node.Left != nil {
            stack = append(stack, Depth{curr.node.Left, curr.depth+1})
        }
    }
    return maxDepth
}

type Depth struct {
    node *TreeNode
    depth int
}

