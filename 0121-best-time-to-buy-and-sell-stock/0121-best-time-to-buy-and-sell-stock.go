func maxProfit(prices []int) int {
    /*
        현재 최저값과 max profit 값을 가지고 있기
        for loop를 돌면서 최저값보다 작으면 최저값 갱신

        max profit return
    */

    var maxProfit int
    var lowest int = prices[0]
    for i:=0; i < len(prices); i++ {
        if prices[i] < lowest {
            lowest = prices[i]
        }

        if prices[i] - lowest > maxProfit {
            maxProfit = prices[i] - lowest
        }
    }

    return maxProfit
}