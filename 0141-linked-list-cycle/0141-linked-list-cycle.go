/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    // use hashMap to record vistied node
    hMap := make(map[*ListNode]bool)
    curr := head
    for curr != nil {
        if _, ok := hMap[curr]; ok {
            return true
        }
        hMap[curr] = true
        curr = curr.Next
    }

    return false
}