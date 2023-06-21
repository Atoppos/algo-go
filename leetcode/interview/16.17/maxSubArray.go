package maxSubArray

//连续数列

func MaxSubArray(nums []int) int {
    max:=nums[0]
    for i:=1;i<len(nums);i++{
        if nums[i]+nums[i-1]>nums[i]{
            nums[i]+=nums[i-1]
        }
        if nums[i]>max{
            max=nums[i]
        }
    }
    return max
}