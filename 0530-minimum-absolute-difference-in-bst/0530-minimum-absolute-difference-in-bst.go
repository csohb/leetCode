/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    min := math.MaxInt
    stack := []*TreeNode{root}
    arr := []int{}

    for len(stack) != 0 {
        pop := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        if pop != nil {
            arr = append(arr, pop.Val)
            stack = append(stack, pop.Right, pop.Left)
        }
    }

    sort.Ints(arr)

    for i:=0; i < len(arr)-1; i++ {
        difference := arr[i+1] - arr[i]
        if difference < min {
            min = difference
        }
    }

    return min
}