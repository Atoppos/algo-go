package arrangeCoins

//排列硬币

func ArrangeCoins(n int) int {
    i:=1
    for ;;i++{
        if n-i>=0{
            n=n-i
        }else{
            break
        }
    }
    return i-1
}