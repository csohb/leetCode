/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
    var answer []int
    if root == nil {
        return answer
    }

    // currLevel의 가장 마지막 노드만 answer append

    currLevel := []*TreeNode{root}

    for len(currLevel) > 0 {
        nextLevel := []*TreeNode{}
        for i:= 0; i < len(currLevel); i++ {
            node := currLevel[i]
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }

            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }

            if i == len(currLevel) -1 {
                answer = append(answer, node.Val)
            }
        } 
        currLevel = nextLevel
    }

    return answer
}