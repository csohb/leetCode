func closeStrings(word1 string, word2 string) bool {
    // operation 1 : all char occurrence is same

    // operation 2 : 

    // if occurrence list or map is identical then it's close string

    if len(word1) != len(word2) {
        return false
    }

    heap1 := &CHeap{}
    heap2 := &CHeap{}
    heap.Init(heap1)
    heap.Init(heap2)

    // store occur
    
    charMap1 := make(map[byte]int)
    charMap2 := make(map[byte]int)
    for i:=0; i < len(word1); i++ {
        charMap1[word1[i]]++
        charMap2[word2[i]]++
    }

    for k, v := range charMap1 {
        if _, ok := charMap2[k]; !ok {
            return false
        } 
        heap.Push(heap1, Close{k, v})
    }

    for k, v := range charMap2 {
        heap.Push(heap2, Close{k, v})
    }


    length := len(charMap1)
    for length > 0 {
        x1 := heap.Pop(heap1).(Close)
        x2 := heap.Pop(heap2).(Close)

        if x1.occr != x2.occr {
            return false
        }


        length--
    }


    return true
}

type Close struct {
    char byte
    occr int
} 

type CHeap []Close

func(h CHeap) Len() int {
    return len(h)
}

func(h CHeap) Less(i, j int) bool {
    return h[i].occr > h[j].occr
}

func(h CHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func(h *CHeap) Push(x interface{}) {
    *h = append(*h, x.(Close))
}

func(h * CHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0:n-1]
    return x
}
