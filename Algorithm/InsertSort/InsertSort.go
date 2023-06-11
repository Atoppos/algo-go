package main

import (
	"fmt"
)
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

func main(){
	var a =[]int{7,5,41,2,1,4}
	InsertSort(a)
	fmt.Println(a)
}