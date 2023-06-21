package arrayaypairsum

import "sort"

//数组拆分

func ArrayPairSum(nums []int) (ans int) {
    sort.Ints(nums)
    for i:=0;i<len(nums);i+=2{
        ans+=nums[i]
    }
    return
}