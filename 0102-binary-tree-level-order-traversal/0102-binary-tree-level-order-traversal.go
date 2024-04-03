/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    var answer [][]int
    if root == nil {
        return answer
    }
    currLevel := make([]*TreeNode, 0)
    currLevel = append(currLevel, root)

    for len(currLevel) > 0 {
        nextLevel := make([]*TreeNode, 0)
        levelArr := []int{}
        for _, node := range currLevel {
            levelArr = append(levelArr, node.Val)
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }

            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }
        answer = append(answer, levelArr)
        currLevel = nextLevel
    }

    return answer 
}