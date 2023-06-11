package main

import "fmt"
//归并排序
func mergesort(s []int)[]int{
	len:=len(s)
	if len==1{
		return s
	}
	m:=len/2
	leftS:=mergesort(s[:m])
	rightS:=mergesort(s[m:])
	return merge(leftS,rightS)
}

func merge(l []int,r []int) []int{
	lLen:=len(l)
	rLen:=len(r)
	res:=make([]int,0)
	lIndex,rIndex:=0,0
	for lIndex<lLen && rIndex<rLen{
		if l[lIndex]>r[rIndex]{
			res=append(res, r[rIndex])
			rIndex++
		}else{
			res=append(res, l[lIndex])
			lIndex++
		}
	}
	if lIndex<lLen{
		res=append(res, l[lIndex:]...)
	}
	if rIndex<rLen{
		res=append(res, r[rIndex:]...)
	}
	return res
}

func main(){
	var s=[]int{3,44,23,1,4,8,5}
	x:=mergesort(s)
	fmt.Println(x)
}