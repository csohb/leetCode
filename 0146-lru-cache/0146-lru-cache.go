// hashmap : store data key - value / when search hashmap -> time complexicty O(1) constant time
// Double Linked List : store the order, but order needs to be updated or the node deleted -> change of the order 
// order should be linked to each other
// head(most recent) (  , , ,  ) tail (least)
// maintain time complexity as O(1)
type Node struct {
    key int
    val int
    prev *Node
    next *Node
}


type LRUCache struct {
    capacity int
    hMap map[int]*Node
    head *Node
    tail *Node
}


func Constructor(capacity int) LRUCache {
    head := &Node{}
    tail := &Node{}
    head.next = tail
    tail.prev = head
    return LRUCache{
        capacity: capacity, 
        hMap: make(map[int]*Node, capacity),
        head: head,
        tail: tail,
    }
}


func (this *LRUCache) Get(key int) int {
    // lookup hMap[key] -> pointing Node, value (return)
    // Get key -> current Node should be go next to Head
    node, ok := this.hMap[key]
    if !ok {
        return -1
    }
    // disconnect from current prev, next 
    this.delete(node)
    // put next to head 
    this.addNextHead(node)

    return node.val
}


func (this *LRUCache) Put(key int, value int)  {
    // check if this.capacity will be over after this put operation
    if node, ok := this.hMap[key]; ok {
        node.val = value
        this.delete(node)
        this.addNextHead(node)
        return
    }

    if len(this.hMap) ==  this.capacity {
        // delete least used Node 
        // least used Node == tail.prev 
        lru := this.tail.prev
        this.delete(lru)
        delete(this.hMap, lru.key)
    } 
    // put new Node
    newNode := &Node{key: key, val: value}
    this.addNextHead(newNode)
    this.hMap[key] = newNode
}

func (this *LRUCache) delete(node *Node) {
    if node == nil || node == this.head || node == this.tail {
        return
    }
    node.prev.next = node.next
    node.next.prev = node.prev

    node.prev = nil
    node.next = nil
    // -> node lost the conneciton between prev & next
}

func (this *LRUCache) addNextHead(node *Node) {
    // put inbetween head // head.next
    old := this.head.next

    this.head.next = node
    node.prev = this.head
    node.next = old

    old.prev = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */