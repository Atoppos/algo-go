package main

import (
	"fmt"
)
//冒泡排序
func Buble(s []int){
	l:=len(s)
	for i:=0;i<l;i++{
		flag:=true
		for j:=0;j<l-i-1;j++{
			if s[j]<s[j+1]{
				s[j],s[j+1]=s[j+1],s[j]
				flag=false
			}
		}
		if flag{
			break
		}
	}
}

func main(){
	var a =[]int{4,5,5,2,9}
	Buble(a)
	fmt.Println(a)

}