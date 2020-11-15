func maxProfit(prices []int) (ans int) {
    for i := 1; i < len(prices); i++ {
        ans += max(0, prices[i]-prices[i-1])   //只要当前价格比第二天价格高，就买进  
    }
    return
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
