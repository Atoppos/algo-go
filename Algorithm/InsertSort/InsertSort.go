package insertsort

//插入排序
func InsertSort(s []int) {
    l := len(s)
    if l < 2 {
        return
    }
    for i:=1;i<l;i++{
        j:=i-1
        value:=s[i]
        for ;j>=0;j--{
            if s[j]>value{
                s[j+1]=s[j]
            }else{
                break
            }
        }
        s[j+1]=value
    }
}
