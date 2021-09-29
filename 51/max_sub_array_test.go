package maxSubArray

// 功能: 给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。
// 思路: 动态规划法

func maxSubArray(nums []int) int {

	dp := make([]int, len(nums))
	dp[0] = nums[0]

	max := nums[0]
	for i:=1; i<len(nums);i++ {
		if dp[i-1] > 0 {
			dp[i] = nums[i] + dp[i-1]
		} else {
			dp[i] = nums[i]
		}

		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}