package removeduplicates

//快慢指针删除有序数组中重复元素
func RemoveDuplicates(s []int) int{
	l:=len(s)
	if l<2{
		return 0
	}
	low:=0
	for fast:=1;fast<l;fast++{
		if s[fast]!=s[low]{
			s[low+1]=s[fast]
			low++
		}
	}
	return low+1
}
