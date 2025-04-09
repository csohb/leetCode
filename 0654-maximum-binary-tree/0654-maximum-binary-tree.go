/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    // get biggest val
    if len(nums) < 1 {
        return nil
    }
    
    
    maxIndex := getMaxIndex(nums)

    prefix := nums[:maxIndex]
    suffix := nums[maxIndex+1:]

    
    root := &TreeNode{Val:nums[maxIndex]}
    root.Left = constructMaximumBinaryTree(prefix)
    root.Right = constructMaximumBinaryTree(suffix)
    return root
    
}

func getMaxIndex(nums []int) int {
    var max int
    var maxIndex int
    for i:=0; i<len(nums); i++ {
        if nums[i] > max {
            max = nums[i]
            maxIndex = i
        } 
    }
    return maxIndex
}