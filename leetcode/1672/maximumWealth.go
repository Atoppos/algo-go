package maximumWealth

//最富有客户资产

func MaximumWealth(accounts [][]int) int {
    if len(accounts)<=1{
        return maxassets(accounts[0])
    }
    var max=0
    for _,p:=range accounts{
        s:=maxassets(p)
        if s>max{
            max=s
        }
    }
    return max
}

func maxassets(assets []int) int{
    var sum=0
    for _,s:=range assets{
        sum+=s
    }
    return sum
}