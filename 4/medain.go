package dedain

// 功能: 给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
// 要求: 时间复杂度为 O(log (m+n))
// 思路: 通过归并的方式将两个数组合并成一个有序的数组, 才取其中位数

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	length1, length2 := len(nums1), len(nums2)
	total := length1 + length2
	nums := make([]int, total)

	lix, liy := 0, 0
	for lix < length1 && liy < length2 {
		if nums1[lix] < nums2[liy] {
			nums[lix+liy] = nums1[lix]
			lix++
		} else {
			nums[lix+liy] = nums2[liy]
			liy++
		}
	}

	if lix < length1 {
		for ; lix < length1; lix++ {
			nums[lix+liy] = nums1[lix]
		}
	}

	if liy < length2 {
		for ; liy < length2; liy++ {
			nums[lix+liy] = nums2[liy]
		}
	}

	if total%2 == 0 {
		return float64(nums[total/2-1]+nums[total/2]) / 2.0
	}
	return float64(nums[total/2])
}

// TODO 通过二分查找来实现
