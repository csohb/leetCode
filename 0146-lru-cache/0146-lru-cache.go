
type LRUCache struct {
    head, tail *listNode
    hMap map[int]*listNode
    capacity int
}

func newLRUCache(head, tail *listNode, cap int) LRUCache {
    return LRUCache{
        head: head,
        tail: tail,
        hMap: make(map[int]*listNode),
        capacity: cap,
    }
}

type listNode struct {
    prev *listNode
    next *listNode
    key, val int
}

func newListNode(key, val int) *listNode {
    return &listNode{
        key: key,
        val: val,
    }
}


func Constructor(capacity int) LRUCache {
    head, tail := newListNode(0,0), newListNode(0,0)
    head.next = tail
    tail.prev = head
    return newLRUCache(head, tail, capacity)
}


func (this *LRUCache) Get(key int) int {
    if n, ok := this.hMap[key]; ok {
        this.remove(n)
        this.insert(n)
        return n.val
    }

    return -1
}


func (this *LRUCache) Put(key int, value int)  {
    if _, ok := this.hMap[key]; ok {
        this.remove(this.hMap[key])
    }

    if len(this.hMap) == this.capacity {
        this.remove(this.tail.prev)
    }

    this.insert(newListNode(key, value))
}


func(this *LRUCache) remove(node *listNode) {
    delete(this.hMap, node.key)
    node.prev.next = node.next
    node.next.prev = node.prev
}

func(this *LRUCache) insert(node *listNode) {
    this.hMap[node.key] = node
    next := this.head.next
    this.head.next = node
    node.prev = this.head
    next.prev = node
    node.next = next
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */