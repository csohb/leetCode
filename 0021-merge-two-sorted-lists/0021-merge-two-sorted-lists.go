/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    start := &ListNode{}
    current := start
    // 일단 추가한다음 뭐가 더 크면 순서 바꾸기?
    for list1 != nil && list2 != nil  {
        if list1.Val < list2.Val {
            current.Next = list1
            list1= list1.Next
        } else {
            current.Next = list2
            list2 = list2.Next
        }
        current = current.Next
     }

     if list1 == nil {
        current.Next = list2
     } else {
        current.Next = list1
     }
    return start.Next 
}