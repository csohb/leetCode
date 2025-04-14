/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    

    hp := &ListHeap{}
    heap.Init(hp)

    for head != nil {
        next := head.Next
        head.Next = nil
        heap.Push(hp, head)
        head = next
    }

    newList := &ListNode{}
    current := newList

    for hp.Len() > 0 {
        node := heap.Pop(hp).(*ListNode)
        current.Next = node
        current = node 
    }

    return newList.Next
}


type ListHeap []*ListNode

func(h ListHeap) Len() int {
    return len(h)
}

func(h ListHeap) Less(i, j int) bool {
    return h[i].Val < h[j].Val
}

func(h ListHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
} 

func(h *ListHeap) Push(x interface{}) {
    *h = append(*h, x.(*ListNode))
} 

func(h *ListHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[:n-1]
    return x
}