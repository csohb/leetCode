func insertionSortList(head *ListNode) *ListNode {
    // 비어있는 리스트거나 노드가 하나뿐인 경우 그대로 반환합니다
    if head == nil || head.Next == nil {
        return head
    }
    
    // 정렬된 리스트의 시작점으로 사용할 더미 노드를 생성합니다
    dummy := &ListNode{} 
    
    // 현재 삽입할 노드를 가리키는 포인터입니다
    curr := head         
    
    // 모든 노드를 순회하며 정렬된 리스트에 삽입합니다
    for curr != nil {
        // 다음에 처리할 노드를 미리 저장해둡니다
        next := curr.Next
        
        // 현재 노드를 삽입할 위치를 찾기 위한 포인터들입니다
        prev := dummy
        p := dummy.Next
        
        // 삽입할 위치를 찾습니다 - 현재 노드의 값보다 작은 값들을 지나갑니다
        for p != nil && p.Val < curr.Val {
            prev = p
            p = p.Next
        }
        
        // 현재 노드를 prev와 p 사이에 삽입합니다
        curr.Next = p
        prev.Next = curr
        
        // 다음 삽입할 노드로 이동합니다
        curr = next
    }
    
    // 더미 노드 다음부터가 실제 정렬된 리스트의 시작점입니다
    return dummy.Next
}