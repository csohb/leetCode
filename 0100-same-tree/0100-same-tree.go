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
     "reflect"
 )
func isSameTree(p *TreeNode, q *TreeNode) bool {
    var answer bool

    if p == nil && q == nil {
        return true
    }

    pList := bfs(p)
    qList := bfs(q)

    if reflect.DeepEqual(pList, qList) {
        answer = true
    } else {
        answer = false
    }

    return answer   
}

func bfs(root *TreeNode) []*TreeNode {
    if root == nil {
        return nil
    }
    queue := make([]*TreeNode, 0)
    queue = append(queue, root)

    nodeList := make([]*TreeNode, 0)
    for len(queue) > 0 {
        start := queue[0]
        queue = queue[1:]
        nodeList = append(nodeList, start)

        if start.Left != nil {
            queue = append(queue, start.Left)
            nodeList = append(nodeList, start.Left)
        }

        if start.Right != nil {
            queue = append(queue, start.Right)
            nodeList = append(nodeList, start.Right)
        }
    }

    return nodeList
}
