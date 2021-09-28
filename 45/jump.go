package jump
// 给你一个非负整数数组 nums ，你最初位于数组的第一个位置。
// 数组中的每个元素代表你在该位置可以跳跃的最大长度。
// 你的目标是使用最少的跳跃次数到达数组的最后一个位置。
// 假设你总是可以到达数组的最后一个位置。
// 思路: 从倒数第二个元素开始反推, 计算每个节点需要的最小步数
// 采用动态规划的方法来实现

func Jump(nums []int) int {
	length := len(nums)
	dp := make([]int, length)

	for i:=length-2; i>=0; i-- {
		min := dp[i+1]
		for j:=1; j<=nums[i] && i+j <= length-1; j++ {
			if dp[i+j] < min {
				min = dp[i+j]
			}
		}
		dp[i] = min + 1
	}

	return dp[0]
}