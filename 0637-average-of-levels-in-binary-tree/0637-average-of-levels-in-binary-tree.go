/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    // how I will check the depth?
    var answer []float64
    currLevel := make([]*TreeNode, 0)
    currLevel = append(currLevel, root)

    for len(currLevel) > 0 {
        nextLevel := make([]*TreeNode, 0)
        var sum int
        for _, node := range currLevel {
            sum += node.Val
            if node.Left != nil {
                nextLevel = append(nextLevel, node.Left)
            }

            if node.Right != nil {
                nextLevel = append(nextLevel, node.Right)
            }
        }
        answer = append(answer, float64(sum) / float64(len(currLevel)))
        currLevel = nextLevel
    }
    return answer 
}
