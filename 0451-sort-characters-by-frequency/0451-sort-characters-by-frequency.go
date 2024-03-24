type CharCounter struct {
    value byte
    count int
}

type MinHeap []CharCounter

func(h MinHeap)Len() int {
    return len(h)
}

func(h MinHeap)Less(i, j int) bool {
    return h[i].count > h[j].count
}

func(h MinHeap)Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func(h *MinHeap)Push(v interface{}) {
    *h = append(*h, v.(CharCounter))
} 

func(h *MinHeap)Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x   
}


func frequencySort(s string) string {
    charMap := make(map[byte]int)

    for i:=0; i < len(s); i++ {
        charMap[s[i]]++
    }

    
    h := &MinHeap{}
    for chr, count := range charMap {
        heap.Push(h, CharCounter{
            value:chr,
            count:count,
        })
    }

    var sb strings.Builder
    for h.Len() > 0 {
        v := heap.Pop(h).(CharCounter)
        for i := 0; i < v.count; i++ {
            sb.WriteString(string(v.value))
        }
    }
    return sb.String()
}