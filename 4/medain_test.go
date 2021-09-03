package dedain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	nums1, nums2 := []int{1, 3}, []int{2, 4}

	medain := FindMedianSortedArrays(nums1, nums2)
	assert.Equal(t, true, 2.5-medain <= 0.0000000000001 && 2.5-medain >= -0.0000000000001)
}
