/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    node := head
    if node == nil {
        return false
    }

    valList := make(map[*ListNode]int)
    
    for ; node.Next != nil; node = node.Next {
        valList[node]++
        if _, ok := valList[node.Next]; ok {
            return true
        }
    }
    return false
}