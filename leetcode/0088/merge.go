package merge

//合并两个有序数组

func Merge(nums1 []int, m int, nums2 []int, n int)  {
    sort:=make([]int,0,m+n)
    p1,p2:=0,0
    for {
        if p1==m{
            sort=append(sort,nums2[p2:]...)
            break
        }
        if p2==n{
            sort=append(sort,nums1[p1:]...)
            break
        }
        if nums1[p1]<nums2[p2]{
            sort=append(sort,nums1[p1])
            p1++
        }else{
            sort=append(sort,nums2[p2])
            p2++
        }
    }
    copy(nums1,sort)
}