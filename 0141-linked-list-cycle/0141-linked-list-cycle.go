/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }

    valList := make(map[*ListNode]int)
    
    for ; head.Next != nil; head = head.Next {
        valList[head]++
        if _, ok := valList[head.Next]; ok {
            return true
        }
    }
    return false
}