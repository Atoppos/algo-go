package finalprices

//商品折扣后的最终价格

func FinalPrices(prices []int) []int {
    if len(prices)<2{
        return prices
    }
    x:=len(prices)
    ans:=make([]int,x)
    stack:=[]int{0}
    for i:=x-1;i>=0;i--{
        now:=prices[i]
        for len(stack)>1&&stack[len(stack)-1]>now{
            stack=stack[:len(stack)-1]
        }
        ans[i]=now-stack[len(stack)-1]
        stack=append(stack,now)
    }
    return ans
}
