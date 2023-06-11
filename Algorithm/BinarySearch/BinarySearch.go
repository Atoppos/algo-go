package main

import "fmt"
//二分搜索
func BinarySearch(target int,s []int) int{
	if len(s)<1{
		return -1
	}
	right:=len(s)-1
	left:=0
	for left<=right{
		mid:=(left+right)/2
		if s[mid]==target{
			return mid
		}
		if s[mid]>target{
			right=mid-1
		}
		if s[mid]<target{
			left=mid+1
		}
	}
	return -1
}

func main(){
	var s=[]int{1,2,3,4,5,6,7,7}
	res:=BinarySearch(1,s)
	fmt.Println(res)
}