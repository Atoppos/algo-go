package fib

//斐波那契数

func Fib(n int) int {
    if n<2{
        return n
    }
    p,q,r:=0,0,1
    for i:=2;i<=n;i++{
        p=q
        q=r
        r=p+q
    }
    return r
}