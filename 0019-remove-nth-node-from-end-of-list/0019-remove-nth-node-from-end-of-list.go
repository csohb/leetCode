/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    // to remove node in List, need to know what is the end
    // from the end, search nth node
    // store prev

    dummy := &ListNode{Next: head}

    left := dummy
    right := head

    for n > 0 && right != nil {
        right = right.Next
        n--
    }


    for right != nil {
        right = right.Next
        left = left.Next
    }


    left.Next = left.Next.Next

    return dummy.Next
}