/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {
        return true
    }

    if p == nil || q == nil {
        return false
    }

    // check if value is identical

    if p.Val != q.Val {
        return false
    }

    queue1 := []*TreeNode{p}
    queue2 := []*TreeNode{q}

    for len(queue1) > 0 {

        level := len(queue1)
        for i:= 0; i < level; i++ {
            curr1 := queue1[0]
            curr2 := queue2[0]

            queue1 = queue1[1:]
            queue2 = queue2[1:]

            if (curr1.Left == nil) != (curr2.Left == nil) {
                return false
            }
        
            if (curr1.Right == nil) != (curr2.Right == nil) {
                return false
            }


            // fmt.Println("curr1.Left:",curr1.Left.Val)
            // //fmt.Println("curr2.Left:",curr2.Left.Val)
            // fmt.Println("curr1.Right:",curr1.Right.Val)
            // fmt.Println("curr2.Right:",curr2.Right.Val)

            // if curr1.Left != curr2.Left || curr1.Right != curr2.Right {
                
            //     return false
            // }

            if curr1.Left != nil {
                if curr1.Left.Val != curr2.Left.Val {
                    return false
                }
                queue1 = append(queue1, curr1.Left)
                queue2 = append(queue2, curr2.Left)
            }
            
            if curr1.Right != nil {
                if curr1.Right.Val != curr2.Right.Val {
                    return false
                }
                queue1 = append(queue1, curr1.Right)
                queue2 = append(queue2, curr2.Right)
            }

        }

    }
    return true
}