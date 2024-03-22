/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */

func copyRandomList(head *Node) *Node {  
    curr := head
    if head == nil {
        return nil
    }

    oldMap := make(map[*Node]*Node)

    for curr != nil {
        oldMap[curr] = &Node{Val:curr.Val}
        curr = curr.Next
    }

    curr = head
    for curr != nil {
        oldMap[curr].Next = oldMap[curr.Next]
        oldMap[curr].Random = oldMap[curr.Random]
        curr = curr.Next
    }

    return oldMap[head]
}