package minmaxgame

//极大极小游戏

func MinMaxGame(nums []int) int {
    if len(nums)==1{
        return nums[0]
    }
    new:=make([]int,len(nums)/2)
    for i:=0;i<len(nums)/2;i++{
        if i%2==0{
            new[i]=min(nums[i*2],nums[i*2+1])
        }else{
            new[i]=max(nums[i*2],nums[i*2+1])
        }
    }
    return MinMaxGame(new)
}

func min(a,b int) int{
    if a>b{
        return b
    }
    return a
}

func max(a,b int) int{
    if a<b{
        return b
    }
    return a
}
