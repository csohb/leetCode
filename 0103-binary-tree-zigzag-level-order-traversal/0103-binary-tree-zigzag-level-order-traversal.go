/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func zigzagLevelOrder(root *TreeNode) [][]int {
    var answer [][]int
    if root == nil {
        return answer
    }
    currLevel := []*TreeNode{root}

    depth := 1
    for len(currLevel) > 0 {
        nextLevel := []*TreeNode{}
        subArr := []int{}
        
        for _, node := range currLevel {
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }

            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }

            if depth % 2 == 1 {
                subArr = append(subArr, node.Val)
            } else {
                subArr = append([]int{node.Val}, subArr...)
            }
        }
        depth++
        answer = append(answer, subArr)
        currLevel = nextLevel
    }

    return answer
}
