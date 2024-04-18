func maxProfit(prices []int) int {
    /*
        첫번째 일단 사기 
        산 가격과 현재 가격 비교
        1. 현재 가격이 더 싸다 -> 산 가격 업데이트 
        2. 현재 가격이 더 비싸다 -> 차액 구하기
            차액 (profit)을 더한 가격을 가지고 있기 (currProfit)
            -> currProfit과 maxProfit을 비교
                -> currProfit이 더 크면 maxProfit에 대입
                    -> buyPrice = 0 으로
                아니면 maxProfit 유지
    */
    var answer int

    for i:= 1; i < len(prices); i++ {
        if prices[i] > prices[i-1] {
            answer += prices[i] - prices[i-1]
        }
    }
    return answer
}
