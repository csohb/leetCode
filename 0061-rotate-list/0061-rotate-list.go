/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil || k == 0 {
        return head
    }

    size := 0
    countHead := head
    for countHead != nil {
        size++
        countHead = countHead.Next
    }

    k %= size

    for k > 0 {
        curr := head
        currVal := curr.Val
        for curr.Next != nil {
            temp := curr.Next.Val 
            curr.Next.Val = currVal
            currVal = temp
            curr = curr.Next
        }
        head.Val = currVal
        k--
    }
    return head
}

