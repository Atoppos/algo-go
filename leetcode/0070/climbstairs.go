package climbstairs

//爬楼梯

func ClimbStairs(n int) int {
    p,q,r:=0,0,1
    for i:=1;i<=n;i++{
        p=q
        q=r
        r=p+q
    }
    return r
}