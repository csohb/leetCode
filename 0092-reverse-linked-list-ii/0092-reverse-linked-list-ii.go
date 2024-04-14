/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
    if left == right || head == nil {
        return head
    }

    dummy := &ListNode{0, head}
    prev := dummy

    for i:=0; i < left - 1; i++ {
        prev = prev.Next
        fmt.Println("prev.Val:", prev.Val)
    }

    current := prev.Next // 2

    for i :=0; i < right - left; i++ {
        nextNode := current.Next // 3
        current.Next = nextNode.Next // current.Next - 4
        nextNode.Next = prev.Next // 4.Next = 2
        prev.Next = nextNode // 3
    }

    return dummy.Next
}