/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    if head == nil || n <= 0 {
        return head
    } 


    // n은 뒤에서 n번째를 의미
    dummy := &ListNode{0, head}
    node := dummy.Next

    var length int
    for node != nil {
        node = node.Next
        length++ //5
    }

    // len = 5, n = 2 3

    fmt.Println("length:",length)

    if length - n == 0 {
        return nil
    }


    current := dummy.Next 

    for i:=1; i <= length; i++ {
        if i == length - n {
            nextNext := current.Next.Next
            current.Next = nextNext
            continue
        } 
        current = current.Next
    }


    return head
}