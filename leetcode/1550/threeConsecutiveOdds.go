package threeConsecutiveOdds

//存在三个奇数的数组

func ThreeConsecutiveOdds(arr []int) bool {
    if len(arr)<3{
        return false
    }
    flag:=false
    for i:=0;i<=len(arr)-3;i++{
        if arr[i]%2!=0{
            if arr[i+1]%2!=0{
                if arr[i+2]%2!=0{
                    flag=true
                }
            }
        }
    }
    return flag
}