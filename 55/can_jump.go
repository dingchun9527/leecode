package canJump

// 功能: 给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 判断你是否能够到达最后一个下标。


func canJump(nums []int) bool {

	length := len(nums)
	dp := make([]bool, length)
	dp[length-1] = true

	for i:=length-2; i>=0; i-- {
		for j:= 0; j<nums[i] && i+j+1 < length; j++ {
			if dp[i+j+1] == true {
				dp[i] = true
			}
		}
	}

	return dp[0]
}