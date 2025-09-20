/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 // head -> Next exists
 // next = nil 
 // curr = curr.Next
 // 
func reverseList(head *ListNode) *ListNode {
    var prev *ListNode // start nil
    curr := head
    for curr != nil {
        next := curr.Next

        curr.Next = prev
        
        prev = curr
        curr = next
    }
    return prev
}