package makefancystring

//删除字符使字符串变好

func MakeFancyString(s string) string{
	result:=make([]byte,0)
    left,right:=0,0
    for right<len(s){
        count:=0
        for right<len(s)&&s[left]==s[right]{
            right++
            count++
            if count<=2{
                result=append(result,s[left])
            }
        }
        left=right
    }
    return string(result)
}
